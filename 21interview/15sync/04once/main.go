package main

import (
	"fmt"
	"sync"
)

func main() {
	var ans int
	var wg sync.WaitGroup

	once := Once{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(func() {
				ans++
			})
		}()
		wg.Done()
	}

	wg.Wait()

	fmt.Println(ans)
}
