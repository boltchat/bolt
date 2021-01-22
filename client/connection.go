package client

import (
	"keesvv/bolt.chat/protocol"
	"net"
)

// Connection TODO
type Connection struct {
	TCPConn *net.TCPConn
	User    protocol.User
}
