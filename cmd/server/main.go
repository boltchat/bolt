package main

import (
	"fmt"
	"net"
)

func logConnection(conn *net.TCPConn) {
	fmt.Printf("%s connected! Say hi.\n", conn.RemoteAddr().String())
}

func handleConnection(conn *net.TCPConn) {
	// Log the connection
	logConnection(conn)

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
