package logging

import (
	"fmt"
	"net"
)

func LogConnection(conn *net.TCPConn) {
	fmt.Printf("%s connected! Say hi.\n", conn.RemoteAddr().String())
}

func LogDisconnect(conn *net.TCPConn) {
	fmt.Printf("%s disconnected.\n", conn.RemoteAddr().String())
}
