package util

import (
	"encoding/json"
	"net"
)

func WriteJson(conn *net.TCPConn, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	conn.Write(b)
	return nil
}
