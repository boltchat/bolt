package main

import (
	"keesvv/go-tcp-chat/internals/message"
	"net"
)

func main() {
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: 3300,
	})

	if err != nil {
		panic(err)
	}

	msg := message.Message{
		Content: "Hi there! ~Kees",
	}

	msg.Send(conn)
	conn.Close()
}
