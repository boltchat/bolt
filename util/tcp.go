package util

import (
	"encoding/json"
	"net"

	"github.com/bolt-chat/server/logging"
)

type ConnPool []*net.TCPConn

func (c *ConnPool) AddToPool(conn *net.TCPConn) {
	*c = append(*c, conn)

	logging.LogDebug(
		"connection added to pool:",
		conn.RemoteAddr().String(),
	)

	logging.LogDebug("pool size:", len(*c))
}

func WriteJson(conn *net.TCPConn, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	conn.Write(b)
}

func Broadcast(conns *ConnPool, data interface{}) {
	for _, conn := range *conns {
		WriteJson(conn, data)
	}
}
