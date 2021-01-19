package handlers

import (
	"bytes"
	"encoding/json"
	"net"

	"keesvv/go-tcp-chat/internals/logging"
	"keesvv/go-tcp-chat/internals/message"
)

/*
HandleConnection handles a TCP connection
during its entire lifespan.
*/
func HandleConnection(conn *net.TCPConn) {
	logging.LogConnection(conn)

	for {
		b := make([]byte, 4096)

		// Wait for and receive new messages
		_, connErr := conn.Read(b)

		if connErr != nil {
			// Broadcast a disconnect message
			logging.LogDisconnect(conn)
			return
		}

		// Trim empty bytes at the end
		b = bytes.TrimRight(b, "\x00")

		msg := &message.Message{}

		// Decode the message
		err := json.Unmarshal(b, msg)

		if err != nil {
			panic(err)
		}

		// Broadcast the message
		msg.Print()
	}
}
