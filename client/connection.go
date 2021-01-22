package client

import (
	"keesvv/go-tcp-chat/protocol"
	"net"
)

// Connection TODO
type Connection struct {
	TCPConn *net.TCPConn
	User    protocol.User
}
