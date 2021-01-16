package logging

import (
	"fmt"
	"net"
)

func LogConnection(conn *net.TCPConn) {
	fmt.Printf("%s connected! Say hi.\n", conn.RemoteAddr().String())
}
