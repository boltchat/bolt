package logging

import (
	"fmt"
	"net"

	"github.com/fatih/color"
)

func logBase(
	level string,
	msg string,
) {
	fmt.Printf("[%s] %s\n", level, msg)
}

func logInfo(msg string) {
	logBase(color.CyanString("INFO"), msg)
}

func LogListener(ip string, port int) {
	logInfo(fmt.Sprintf("Server listening on %s:%v.", ip, port))
}

func LogConnection(conn *net.TCPConn) {
	logInfo(fmt.Sprintf("%s connected! Say hi.", conn.RemoteAddr().String()))
}

func LogDisconnect(conn *net.TCPConn) {
	logInfo(fmt.Sprintf("%s disconnected.", conn.RemoteAddr().String()))
}
