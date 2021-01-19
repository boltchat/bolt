package client

import (
	"encoding/json"
	"keesvv/go-tcp-chat/internals/events"
	"keesvv/go-tcp-chat/internals/message"
	"net"
	"time"
)

/*
SendMessage sends a message to an established
TCP connection.
*/
func SendMessage(m *message.Message, conn *net.TCPConn) error {
	m.SentAt = time.Now().Unix()
	evt := events.NewMessageEvent(m)

	b, err := json.Marshal(*evt)
	if err != nil {
		return err
	}

	conn.Write(b)
	return nil
}
