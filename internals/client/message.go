package client

import (
	"keesvv/go-tcp-chat/internals/protocol"
	"keesvv/go-tcp-chat/internals/protocol/events"
	"keesvv/go-tcp-chat/internals/util"
	"time"
)

/*
SendMessage sends a message to an established
TCP connection.
*/
func (c *Connection) SendMessage(m *protocol.Message) error {
	m.SentAt = time.Now().Unix()
	util.WriteJson(c.TCPConn, *events.NewMessageEvent(m))
	return nil
}
