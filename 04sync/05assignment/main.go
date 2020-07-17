package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var once sync.Once
var wg sync.WaitGroup

func main() {
	wg.Add(25)
	// 启动一个goroutine-随机生成int64类型的值，发送到jobChan中
	ch1, ch2 := make(chan int64), make(chan int64)

	go generate(ch1)

	// 启动24个goroutine，从jobChan中获取值，计算个位数的和，发送到resultChan中
	for i := 0; i < 24; i++ {
		go calculate(ch1, ch2)
	}

	// fmt.Println()
	// 从主goroutine中取出结果并打印到终端
	for c := range ch2 {
		fmt.Println(c)
	}
	wg.Wait()
}

func generate(ch1 chan<- int64) {
	defer close(ch1)
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch1 <- rand.Int63n(1000)
	}
}

func calculate(ch1 <-chan int64, ch2 chan<- int64) {
	// defer func() {
	// 	once.Do(func() { close(ch2) })
	// }()
	defer wg.Done()

	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ans := addDigits(x)

		ch2 <- ans
	}
}

func addDigits(num int64) (ans int64) {
	// times := int64(1)
	for num > 0 {
		carry := num % 10
		num /= 10
		// ans += carry * times
		// times *= 10
		ans += carry
	}
	return
}
