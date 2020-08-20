package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 场景：在一个高并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。
// 每个IP三分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出 success:100

type Ban struct {
	visitIPs map[string]time.Time
	mu       sync.Mutex
}

func NewBan(ctx context.Context) *Ban {
	o := &Ban{visitIPs: make(map[string]time.Time)}

	go func() {
		tick := time.Tick(time.Minute)

		for {
			select {
			case <-ctx.Done():
				return
			case <-tick:
				o.mu.Lock()
				for ip, duration := range o.visitIPs {
					if time.Since(duration) > time.Minute*3 {
						delete(o.visitIPs, ip)
					}
				}
				o.mu.Unlock()
			}
		}
	}()

	return o
}

func (o *Ban) visit(ip string) bool {
	o.mu.Lock()
	defer o.mu.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}

func main() {
	var success int64
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	ban := NewBan(ctx)
	defer cancel()

	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					//success++
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}
	}

	wg.Wait()
	fmt.Println("success:", success)
}
