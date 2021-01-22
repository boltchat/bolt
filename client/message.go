package client

import (
	"keesvv/bolt.chat/protocol"
	"keesvv/bolt.chat/protocol/events"
	"keesvv/bolt.chat/util"
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
