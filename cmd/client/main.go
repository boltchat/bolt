package main

import (
	"bufio"
	"keesvv/go-tcp-chat/internals/client"
	"os"
)

func main() {
	conn, err := client.Connect("Kees", "127.0.0.1", 3300)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		client.Prompt(reader, conn)
	}
}
