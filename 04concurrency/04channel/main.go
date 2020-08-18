package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	wg.Add(2)

	go f1(ch1)
	go f2(ch1, ch2)
	go f2(ch1, ch2)

	for v := range ch2 {
		fmt.Println(v)
	}
	wg.Wait()
}

// 往ch1中写入
func f1(c chan int) {
	defer wg.Done()
	defer close(c) // 如果不关闭，会导致死锁
	for i := 0; i < 100; i++ {
		c <- i
	}
}

// 获取ch1中的值，写入ch2
func f2(ch1, ch2 chan int) {
	defer wg.Done()
	defer func() {
		once.Do(func() { close(ch2) })
	}() // 如果不关闭会导致死锁
	for {
		if i, ok := <-ch1; ok {
			ch2 <- i * i
		} else {
			break
		}
	}
}
