package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

func main() {
	// 开启多个goroutine同时对一个FD进行写入
	goNum := 1000
	var wg sync.WaitGroup

	file, err := os.OpenFile("in.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}

	buf := bufio.NewWriter(file)

	wg.Add(goNum)
	for i := 0; i < goNum; i++ {
		go func() {
			defer func() {
				wg.Done()
				buf.Flush()
			}()
			s := []byte("Hello\n")
			_, err := buf.Write(s)
			if err != nil {
				log.Println(err)
			}
		}()
	}
	wg.Wait()
}
