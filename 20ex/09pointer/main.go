package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}

	newSlice := append(slice, 5)

	fmt.Printf("%p, %p", &slice, &newSlice)
}
