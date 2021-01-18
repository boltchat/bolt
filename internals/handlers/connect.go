package handlers

import (
	"bytes"
	"encoding/json"
	"net"

	"keesvv/go-tcp-chat/internals/logging"
	"keesvv/go-tcp-chat/internals/message"
)

func HandleConnection(conn *net.TCPConn) {
	// TODO: clean up this absolute mess
	logging.LogConnection(conn)

	rawBytes := make([]byte, 4096)
	_, connErr := conn.Read(rawBytes)

	if connErr != nil {
		return
	}

	b := bytes.TrimRight(rawBytes, "\x00")

	msg := &message.Message{}
	err := json.Unmarshal(b, msg)

	if err != nil {
		panic(err)
	}

	msg.Print()
}
