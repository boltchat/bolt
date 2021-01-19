package client

import (
	"keesvv/go-tcp-chat/internals/user"
	"net"
)

// Connection TODO
type Connection struct {
	TCPConn *net.TCPConn
	User    user.User
}
