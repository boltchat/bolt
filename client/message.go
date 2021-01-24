package client

import (
	"time"

	"github.com/bolt-chat/protocol"
	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/util"
)

/*
SendMessage sends a message to an established
TCP connection.
*/
func (c *Connection) SendMessage(m *protocol.Message) error {
	m.SentAt = time.Now().Unix()
	util.WriteJson(c.Conn, *events.NewMessageEvent(m))
	return nil
}
