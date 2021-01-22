package client

import (
	"net"

	"github.com/keesvv/bolt.chat/protocol"
)

// Connection TODO
type Connection struct {
	TCPConn *net.TCPConn
	User    protocol.User
}
