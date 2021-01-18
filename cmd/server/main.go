package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"

	"keesvv/go-tcp-chat/internals/logging"
	"keesvv/go-tcp-chat/internals/message"
)

func handleConnection(conn *net.TCPConn) {
	// TODO: clean up this absolute mess
	logging.LogConnection(conn)

	rawBytes := make([]byte, 4096)
	conn.Read(rawBytes)

	b := bytes.TrimRight(rawBytes, "\x00")

	msg := &message.Message{}
	err := json.Unmarshal(b, msg)

	if err != nil {
		panic(err)
	}

	fmt.Println(msg)
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
