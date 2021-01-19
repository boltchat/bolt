package main

import (
	"bufio"
	"keesvv/go-tcp-chat/internals/client"
	"os"
)

func main() {
	conn, err := client.Connect(client.Options{
		Hostname: "127.0.0.1",
		Port:     3300,
		Nickname: "keesvv",
	})

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		client.Prompt(reader, conn)
	}
}
