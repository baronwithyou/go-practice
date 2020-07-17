package main

import "fmt"

// 闭包
func main() {
	f1(f3(100, 200))
}

func f1(f func()) {
	f()
}

func f2(x, y int) {
	fmt.Println(x + y)
}

// 目的：将f2传入到f1中
func f3(x, y int) func() {
	return func() {
		f2(x, y)
	}
}