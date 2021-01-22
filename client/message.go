package client

import (
	"time"

	"github.com/keesvv/bolt.chat/protocol"
	"github.com/keesvv/bolt.chat/protocol/events"
	"github.com/keesvv/bolt.chat/util"
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
