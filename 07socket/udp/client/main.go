package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0,0,0,0),
		Port: 1235,
	})
	if err != nil {
		log.Printf("fail to dial server, err: %v\n", err)
		return
	}

	defer socket.Close()
	sendData := []byte("Hello Server")
	_, err = socket.Write(sendData)
	if err != nil {
		log.Printf("fail to send message, err: %v", err)
		return
	}

	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		log.Printf("fail to read message, err: %v", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr,n)
}
