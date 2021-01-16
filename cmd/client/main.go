package main

import (
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

	conn.Write([]byte("Hi there! ~Kees"))
	conn.Close()
}
