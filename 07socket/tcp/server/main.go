package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed, err: %s\n", err)
		return
	}

	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err: %s\n", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)

		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from client failed. err: %s\n", err)
			break
		}

		recvStr := string(buf[:n])
		fmt.Printf("receive data from client: %s\n", recvStr)
		conn.Write([]byte(recvStr))
	}
}
