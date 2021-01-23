package util

import (
	"encoding/json"
	"net"
)

func WriteJson(conn net.Conn, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	conn.Write(b)
}

func Broadcast(conns []net.Conn, data interface{}) {
	for _, conn := range conns {
		WriteJson(conn, data)
	}
}
