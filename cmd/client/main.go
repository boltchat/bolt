package main

import (
	"bufio"
	"keesvv/go-tcp-chat/internals/client"
	"net"
	"os"
)

func main() {
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: 3300,
	})

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		client.Prompt(reader, conn)
	}
}
