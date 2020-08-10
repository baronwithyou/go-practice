package main

import "sync"

/**
问题描述:

使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：

12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/

func main() {

	letter, number := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				print(i)
				i++
				print(i)
				i++
				letter <- true
			}
		}
	}()
	wg.Add(1)
	// 需要注意的是这里需要传指针
	go func(wg *sync.WaitGroup) {
		i := 0
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

		for {
			select {
			case <-letter:
				if i >= len(str) {
					wg.Done()
					return
				}
				print(str[i : i+1])
				i++
				print(str[i : i+1])
				i++
				number <- true
			}
		}
	}(&wg)

	number <- true
	wg.Wait()
}
