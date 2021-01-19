package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"

	"keesvv/go-tcp-chat/internals/events"
	"keesvv/go-tcp-chat/internals/logging"
)

/*
HandleConnection handles a TCP connection
during its entire lifespan.
*/
func HandleConnection(conn *net.TCPConn) {
	logging.LogConnection(conn)

	for {
		b := make([]byte, 4096)

		// Wait for and receive incoming events
		_, connErr := conn.Read(b)

		if connErr != nil {
			// Broadcast a disconnect message
			logging.LogDisconnect(conn)
			return
		}

		// Trim empty bytes at the end
		b = bytes.TrimRight(b, "\x00")

		evt := &events.BaseEvent{}

		// Decode the event
		err := json.Unmarshal(b, evt)

		if err != nil {
			panic(err)
		}

		if evt.Event.Type == events.MessageType {
			fmt.Printf("this is a message event!\n")
		}
	}
}
