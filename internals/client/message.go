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
func (c *Connection) SendMessage(m *message.Message) error {
	m.SentAt = time.Now().Unix()
	evt := events.NewMessageEvent(m)

	b, err := json.Marshal(*evt)
	if err != nil {
		return err
	}

	c.TCPConn.Write(b)
	return nil
}
