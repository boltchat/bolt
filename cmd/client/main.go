package main

import (
	"bufio"
	"fmt"
	"keesvv/go-tcp-chat/internals/message"
	"keesvv/go-tcp-chat/internals/user"
	"net"
	"os"
	"strings"
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
		fmt.Print("> ")
		content, readError := reader.ReadString('\n')
		content = strings.TrimSpace(content)

		if readError != nil {
			panic(readError)
		}

		msg := message.Message{
			Content: content,
			User: &user.User{
				Nickname: "Kees",
			},
		}

		msg.Send(conn)
	}
}
