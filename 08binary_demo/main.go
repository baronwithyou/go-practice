package main

import "fmt"

// 二进制实际用途

// 111
// 左边的1表示吃饭 100
// 中间的1表示睡觉 010
// 右边的1表示打豆豆 001

const (
	eat   int = 4
	sleep int = 2
	beat  int = 1
)

func main() {
	f(eat | sleep | beat)
}

func f(arg int) {
	fmt.Printf("%b\n", arg)
}
