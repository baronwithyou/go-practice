package main

import "fmt"

func main() {
	// 数组的三种初始化方式
	a1 := [5]int{1, 2}
	fmt.Printf("Type: %T Value: %#v\n", a1, a1)

	a2 := [...]string{"hello", "world"}
	fmt.Printf("Type: %T Value: %#v\n", a2, a2)

	a3 := [5]bool{1: true, 4: true}
	fmt.Printf("Type: %T Value: %#v\n", a3, a3)

	// 定义多维数组
	// [[1,2], [3,4], [5,6]]
	a4 := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Printf("Type: %T Value: %#v\n", a4, a4)
}
