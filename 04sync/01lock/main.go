package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex
var set = make(map[int]bool, 0)

func main() {
	for i := 0; i < 10; i++ {
		go printOnce(100)
	}
	time.Sleep(time.Second)
}

func printOnce(num int) {
	m.Lock()
	defer m.Unlock()
	if _, ok := set[num]; !ok {
		fmt.Println(num)
	}

	set[num] = true
	// m.Unlock()
}
