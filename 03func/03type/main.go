package main

import "fmt"

// CalculateType ...
type CalculateType func(int, int) int

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

// Calculate ...
func Calculate(a, b int, f CalculateType) int {
	return f(a, b)
}

func main() {
	a, b := 2, 3

	fmt.Println(Calculate(a, b, add))
	fmt.Println(Calculate(a, b, mul))
}
