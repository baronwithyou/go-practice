package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(num)
		}(i, c)
	}

	if WaitTimeout(&wg, time.Second*5) {
		close(c)
		fmt.Println("timeout exit")
	}
	time.Sleep(time.Second * 10)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false

	timeoutChan := time.After(timeout)
	doneChan := make(chan bool)

	go func() {
		defer func() { doneChan <- true }()

		wg.Wait()
	}()

	select {
	case <-timeoutChan:
		return true
	case <-doneChan:
		return false
	}
}

// 这种方式更简洁
func WaitTimeout2(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false

	ch := make(chan bool)

	time.AfterFunc(timeout, func() {
		ch <- true
	})

	go func() {
		wg.Wait()

		ch <- false
	}()

	return <-ch
}
