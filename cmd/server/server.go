package main

import (
	"fmt"
	"os"

	"github.com/bolt-chat/server"
)

func main() {
	ipv4Bind := os.Getenv("BIND_IPV4")
	ipv6Bind := os.Getenv("BIND_IPV6")

	if ipv4Bind == "" {
		// Set default IPv4 bind to loopback address
		ipv4Bind = "127.0.0.1"
	}

	if ipv6Bind == "" {
		// Set default IPv6 bind to loopback address
		ipv6Bind = "::1"
	}

	listener := server.Listener{
		Bind: []server.Bind{
			{Address: ipv4Bind, Proto: "tcp4"},
			{Address: ipv6Bind, Proto: "tcp6"},
		},
		Port: 3300,
	}

	err := listener.Listen()
	if err != nil {
		panic(err)
	}

	// Exit on return
	fmt.Scanln()
}
