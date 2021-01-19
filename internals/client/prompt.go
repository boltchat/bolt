package client

import (
	"bufio"
	"fmt"
	"keesvv/go-tcp-chat/internals/message"
	"keesvv/go-tcp-chat/internals/user"
	"net"
	"strings"
)

/*
Prompt prompts the user for sending messages.
*/
func Prompt(r *bufio.Reader, conn *net.TCPConn) {
	fmt.Print("> ")
	content, readError := r.ReadString('\n')
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

	err := SendMessage(&msg, conn)
	if err != nil {
		panic(err)
	}
}
