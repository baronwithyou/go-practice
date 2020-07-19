package main

import (
	"log"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.IPv4(0,0,0,0),
		Port: 1235,
	})
	if err != nil {
		log.Printf("listen failed, err: %v\n", err)
		return
	}

	defer listen.Close()
	for {
		var data [1024]byte
		// TCP的话需要接受建立连接
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			log.Printf("fail to read msg, err: %v\n", err)
			return
		}

		log.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)

		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			log.Printf("fail to write msg, err: %v\n", err)
			continue
		}
	}
}
