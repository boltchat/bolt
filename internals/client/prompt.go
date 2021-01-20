package client

import (
	"bufio"
	"fmt"
	"keesvv/go-tcp-chat/internals/message"
	"strings"

	"github.com/fatih/color"
)

/*
Prompt prompts the user for sending messages.
*/
func Prompt(r *bufio.Reader, conn *Connection) {
	fmt.Print(color.CyanString("> "))
	content, readError := r.ReadString('\n')
	content = strings.TrimSpace(content)

	if content == "" {
		return
	}

	if readError != nil {
		panic(readError)
	}

	msg := message.Message{
		Content: content,
		User:    &conn.User,
	}

	err := conn.SendMessage(&msg)
	if err != nil {
		panic(err)
	}
}
