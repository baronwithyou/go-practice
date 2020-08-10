package main

import "sync"

/**
问题描述:

使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：

12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/

var wg sync.WaitGroup

func main() {

	numChan := make(chan struct{}, 1)
	charChan := make(chan struct{}, 1)

	defer close(numChan)
	defer close(charChan)

	printNum(numChan, charChan)
	printChar(numChan, charChan)

	numChan <- struct{}{}

	wg.Wait()
}

func printNum(numChan chan struct{}, charChan chan struct{}) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 29; i++ {
			select {
			case <-numChan:
				print(i)
				if i%2 == 0 {
					charChan <- struct{}{}
				} else {
					numChan <- struct{}{}
				}
			}
		}
	}()
}

func printChar(numChan chan struct{}, charChan chan struct{}) {
	wg.Add(1)

	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	go func() {
		defer wg.Done()

		for i := 0; i < len(str); i++ {
			select {
			case <-charChan:
				print(str[i : i+1])
				if (i+1)%2 == 0 {
					numChan <- struct{}{}
				} else {
					charChan <- struct{}{}
				}
			}
		}
	}()
}
