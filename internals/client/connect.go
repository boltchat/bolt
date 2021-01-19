package client

import (
	"encoding/json"
	"keesvv/go-tcp-chat/internals/events"
	"keesvv/go-tcp-chat/internals/user"
	"net"
)

func Connect(ip string, port int) (*net.TCPConn, error) {
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	})

	if err != nil {
		return nil, err
	}

	evt := events.NewJoinEvent(&user.User{
		Nickname: "Kees",
	})

	b, jsonErr := json.Marshal(evt)
	if jsonErr != nil {
		return nil, jsonErr
	}

	conn.Write(b)
	return conn, nil
}
