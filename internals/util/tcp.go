package util

import (
	"encoding/json"
	"net"
)

func WriteJson(conn *net.TCPConn, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	conn.Write(b)
}

func Broadcast(conns []*net.TCPConn, data interface{}) {
	for _, conn := range conns {
		WriteJson(conn, data)
	}
}
