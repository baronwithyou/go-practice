package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := newMap()

	var wg sync.WaitGroup
	key := "hello"
	m.Out(key, 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			fmt.Println(m.Rd(key, time.Microsecond*2))
		}(i)
	}

	wg.Wait()
}

type sp interface {
	Out(key string, val interface{})                  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type Map struct {
	mu sync.RWMutex

	data map[string]interface{}
}

func newMap() *Map {
	m := &Map{}

	m.data = make(map[string]interface{}, 0)

	return m
}

func (m *Map) Out(key string, val interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[key] = val
}

func (m *Map) Rd(key string, timeout time.Duration) interface{} {
	boom := time.After(timeout)

	time.Sleep(time.Microsecond)
	select {
	case <-boom:
		fmt.Println("timeout")
		return 0
	default:
		m.mu.Lock()
		defer m.mu.Unlock()

		return m.data[key]
	}
}
