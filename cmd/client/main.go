package main

import (
	"keesvv/go-tcp-chat/internals/message"
	"keesvv/go-tcp-chat/internals/user"
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
		Content: "Hi there!",
		User: &user.User{
			Nickname: "Kees",
		},
	}

	msg.Send(conn)
	conn.Close()
}
