package client

import (
	"keesvv/go-tcp-chat/internals/protocol"
	"net"
)

// Connection TODO
type Connection struct {
	TCPConn *net.TCPConn
	User    protocol.User
}
