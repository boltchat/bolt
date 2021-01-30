package client

import (
	"github.com/bolt-chat/protocol"
	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/util"
)

/*
SendMessage sends a message to an established
TCP connection.
*/
func (c *Client) SendMessage(m *protocol.Message) error {
	util.WriteJson(c.Conn, *events.NewMessageEvent(m))
	return nil
}
