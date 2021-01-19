package client

import (
	"encoding/json"
	"keesvv/go-tcp-chat/internals/events"
	"keesvv/go-tcp-chat/internals/user"
	"net"
)

func Connect(nick string, ip string, port int) (*Connection, error) {
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	})

	if err != nil {
		return &Connection{}, err
	}

	user := &user.User{
		Nickname: nick,
	}

	evt := events.NewJoinEvent(user)
	b, jsonErr := json.Marshal(evt)

	if jsonErr != nil {
		return &Connection{}, jsonErr
	}

	conn.Write(b)

	return &Connection{
		TCPConn: conn,
		User:    *user,
	}, nil
}
