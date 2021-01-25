package util

import (
	"encoding/json"
	"net"
)

type ConnPool []*net.TCPConn

func (c *ConnPool) AddToPool(conn *net.TCPConn) {
	*c = append(*c, conn)
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
