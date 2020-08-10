package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// 实现一个聊天室
var (
	connections = make(map[string]net.Conn, 0)
)

type Message struct {
	name    string
	content string
	time    time.Time
}

func main() {
	listen, err := net.Listen("tcp", "192.168.31.110:1234")
	if err != nil {
		log.Printf("fail to listen the address, err: %v\n", err)
		return
	}

	defer listen.Close()

	messageChan := make(chan *Message)
	defer close(messageChan)

	// 将消息广播出去
	go sendMessage(messageChan)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("fail to accept connection, err: %v\n", err)
			continue
		}

		go func() {
			name, err := initName(conn)
			if err != nil {
				log.Printf("fail to init client name, err: %v\n", err)
				return
			}

			log.Printf("%s enter the chatting room\n", name)

			connections[name] = conn

			go process(name, conn, messageChan)
		}()
	}

}

func initName(conn net.Conn) (string, error) {
	for {
		conn.Write([]byte("write your name:"))

		reader := bufio.NewReader(conn)

		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			// log.Printf("fail to receive message, err: %v\n", err)
			return "", fmt.Errorf("fail to receive message, err: %v\n", err)
		}
		if n == 0 {
			conn.Write([]byte("the name is empty"))
			continue
		}

		name := string(buf[:n])

		if _, ok := connections[name]; ok {
			conn.Write([]byte("the name is exists"))
			continue
		}
		return name, nil
	}
}

func process(name string, conn net.Conn, ch chan<- *Message) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)

		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			log.Printf("fail to receive message, err: %v\n", err)
			return
		}

		msg := &Message{
			name:    name,
			content: string(buf[:n]),
		}
		ch <- msg
	}
}

func sendMessage(ch <-chan *Message) {
	for s := range ch {
		for _, conn := range connections {
			conn.Write(formatMessage(s))
		}
	}
}

func formatMessage(m *Message) []byte {
	s := fmt.Sprintf("%s: %s", m.name, m.content)

	return []byte(s)
}
