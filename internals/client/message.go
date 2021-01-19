package client

import (
	"encoding/json"
	"keesvv/go-tcp-chat/internals/events"
	"keesvv/go-tcp-chat/internals/message"
	"time"
)

/*
SendMessage sends a message to an established
TCP connection.
*/
func SendMessage(m *message.Message, conn *Connection) error {
	m.SentAt = time.Now().Unix()
	evt := events.NewMessageEvent(m)

	b, err := json.Marshal(*evt)
	if err != nil {
		return err
	}

	conn.TCPConn.Write(b)
	return nil
}
