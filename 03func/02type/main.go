package main

import "fmt"

// CalculateType a type of calculation
type CalculateType func(int, int) int

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func main() {
	a := CalculateType(add)
	b := CalculateType(sub)

	fmt.Println(a(1, 2))
	fmt.Println(b(123, 2))
}
