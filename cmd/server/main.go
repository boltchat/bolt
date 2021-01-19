package main

import (
	"keesvv/go-tcp-chat/internals/server"
)

func main() {
	listener := server.Listener{
		IP:   "127.0.0.1",
		Port: 3300,
	}

	err := listener.Listen()
	if err != nil {
		panic(err)
	}
}
