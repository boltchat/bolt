package client

import (
	"keesvv/go-tcp-chat/internals/events"
	"keesvv/go-tcp-chat/internals/message"
	"keesvv/go-tcp-chat/internals/util"
	"time"
)

/*
SendMessage sends a message to an established
TCP connection.
*/
func (c *Connection) SendMessage(m *message.Message) error {
	m.SentAt = time.Now().Unix()
	util.WriteJson(c.TCPConn, *events.NewMessageEvent(m))
	return nil
}
