package handlers

import (
	"bytes"
	"encoding/json"
	"net"

	"keesvv/go-tcp-chat/protocol/events"
	"keesvv/go-tcp-chat/server/logging"
	"keesvv/go-tcp-chat/util"
)

/*
HandleConnection handles a TCP connection
during its entire lifespan.
*/
func HandleConnection(conns []*net.TCPConn, conn *net.TCPConn) {
	for {
		// a := server.Listener{}
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
			util.WriteJson(conn, *events.NewErrorEvent("invalid_format"))
			conn.Close()
			return
		}

		switch evt.Event.Type {
		case events.MessageType:
			msgEvt := &events.MessageEvent{}
			json.Unmarshal(b, msgEvt)
			util.Broadcast(conns, msgEvt) // TODO: mutate and write
		case events.JoinType:
			joinEvt := &events.JoinEvent{}
			json.Unmarshal(b, joinEvt)
			logging.LogConnection(conn) // TODO

			util.WriteJson(conn, *events.NewMotdEvent("This is the message of the day!")) // TODO
		default:
			// TODO: event not understood
		}
	}
}
