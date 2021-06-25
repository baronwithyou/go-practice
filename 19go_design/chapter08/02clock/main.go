package main

import (
	"log"
	"net"
	"io"
	"time"
)


func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// 当有请求来的时候，会使用主线程进行处理，所以会阻塞其他请求
		// handleConn(conn)

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Println("client close the connection")
			return
		}

		time.Sleep(1 * time.Second)
	}
}