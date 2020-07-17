package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:20000")
	if err != nil {
		fmt.Printf("dial failed: err: %s", err)
		return
	}

	for i := 0; i < 20; i++ {
		msg := "Helo"
		conn.Write([]byte(msg))
	}
}
