package message

import (
	"encoding/json"
	"fmt"
	"keesvv/go-tcp-chat/internals/user"
	"net"
	"time"
)

/*
Message represents a message that is
either transmitted or stored locally.
*/
type Message struct {
	SentAt  int64      `json:"sent"`
	Content string     `json:"content"`
	User    *user.User `json:"user"`
}

/*
Send sends the message to an established
TCP connection.
*/
func (m *Message) Send(conn *net.TCPConn) error {
	m.SentAt = time.Now().Unix()

	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	conn.Write(b)
	return nil
}

/*
Print prints a message to stdout.
*/
func (m *Message) Print() {
	sentAt := time.Unix(m.SentAt, 0)

	fmt.Printf(
		"[%v] <%s> %s\n",
		sentAt.Format(time.Stamp),
		m.User.Nickname,
		m.Content,
	)
}
