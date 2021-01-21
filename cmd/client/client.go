package main

import (
	"bufio"
	"keesvv/go-tcp-chat/internals/client"
	"os"
)

func main() {
	conn, err := client.Connect(client.Options{
		Hostname: "localhost",
		Port:     3300,
		Nickname: "keesvv",
	})

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	go conn.HandleEvents()

	client.Prompt(reader, conn)
}
