package client

import (
	"crypto/tls"

	"github.com/bolt-chat/protocol"
)

// Connection TODO
type Connection struct {
	Conn *tls.Conn
	User protocol.User
}
