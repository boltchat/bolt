package logging

import (
	"fmt"
	"net"
	"time"

	"github.com/fatih/color"
)

func logBase(
	level string,
	msg string,
) {
	fmt.Printf("%s [%s] %s\n", color.HiBlackString(time.Now().Format("15:04:05")), level, msg)
}

func LogInfo(msg string) {
	logBase(color.CyanString("INFO"), msg)
}

func LogListener(ip string, port int) {
	LogInfo(fmt.Sprintf("Server listening on %s:%v.", ip, port))
}

func LogConnection(conn *net.TCPConn) {
	LogInfo(fmt.Sprintf("%s connected! Say hi.", conn.RemoteAddr().String()))
}

func LogDisconnect(conn *net.TCPConn) {
	LogInfo(fmt.Sprintf("%s disconnected.", conn.RemoteAddr().String()))
}
