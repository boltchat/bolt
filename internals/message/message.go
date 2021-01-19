package message

import (
	"fmt"
	"keesvv/go-tcp-chat/internals/user"
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
