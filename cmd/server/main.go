package main

import (
	"fmt"
	"net"

	"keesvv/go-tcp-chat/internals/logging"
)

func handleConnection(conn *net.TCPConn) {
	// Log the connection
	logging.LogConnection(conn)

	b := make([]byte, 512)
	conn.Read(b)
	fmt.Println(string(b))
}

func main() {
	l, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: 3300,
	})

	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.AcceptTCP()

		if err != nil {
			panic(err)
		}

		// Accept new connection
		go handleConnection(conn)
	}
}
