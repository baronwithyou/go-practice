package main

import (
	"fmt"
	"math/rand"
)

// var wga sync.WaitGroup

func main() {
	c1 := randomInput(2)
	c2 := randomInput(2)
	// c3 := randomInput(30)
	// c4 := randomInput(40)

	// wga.Add(1)
	p := merge(c1, c2)

	// wga.Wait()
	var count int
	for v := range p {
		count++
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
	fmt.Printf("count of int: %d\n", count)
}

func randomInput(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(500)
		}

		close(out)
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {

	out := make(chan int)

	// var wg sync.WaitGroup
	// output := func(c <-chan int) {
	// 	defer wg.Done()
	// 	for v := range c {
	// 		out <- v
	// 	}
	// }

	// for _, c := range cs {
	// 	wg.Add(1)
	// 	go output(c)
	// }

	// go func() {
	// 	defer wga.Done()
	// 	defer close(out)
	// 	wg.Wait()
	// }()

	go func() {
		// defer wga.Done()
		defer close(out)
		for _, c := range cs {
			for v := range c {
				out <- v
			}
		}
	}()

	return out
}
