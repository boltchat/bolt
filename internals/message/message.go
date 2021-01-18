package message

import (
	"encoding/json"
	"keesvv/go-tcp-chat/internals/user"
	"net"
)

/*
Message represents a message that is
either transmitted or stored locally.
*/
type Message struct {
	Content string     `json:"content"`
	User    *user.User `json:"user"`
}

/*
Send sends the message to an established
TCP connection.
*/
func (m *Message) Send(conn *net.TCPConn) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	conn.Write(b)
	return nil
}
