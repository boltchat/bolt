package logging

import (
	"fmt"
	"net"
)

func LogListener(ip string, port int) {
	fmt.Printf("Server listening on %s:%v.\n", ip, port)
}

func LogConnection(conn *net.TCPConn) {
	fmt.Printf("%s connected! Say hi.\n", conn.RemoteAddr().String())
}

func LogDisconnect(conn *net.TCPConn) {
	fmt.Printf("%s disconnected.\n", conn.RemoteAddr().String())
}
