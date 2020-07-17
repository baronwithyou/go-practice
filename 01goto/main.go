package main

import "fmt"

func main() {
	breakTest()
}

func gotoTest() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 2; j++ {
			if i == 5 {
				goto breakTag
			}

			fmt.Printf("%d\t%d\n", i, j)
		}
	}

breakTag:
	fmt.Println("over")
}

func breakTest() {
BREAK:
	for i := 0; i < 10; i++ {
		if i == 5 {
			break BREAK
		}

		fmt.Println(i)
	}
}
