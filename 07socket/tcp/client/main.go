package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.31.110:20000")
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}

	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}

		_, err = conn.Write([]byte(inputInfo))
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("recv failed, err: %s", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
