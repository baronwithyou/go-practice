package main

import (
	"fmt"
	"time"
)

func main() {
	abc := make(chan int, 20)
	go func() {
		for i := 0; i < 10; i++ {
			abc <- i
		}
	}()

	go func() {
		for {
			a, ok := <-abc
			if !ok {
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(abc)
	fmt.Println("close")
	time.Sleep(time.Second * 100)
}
