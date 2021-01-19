package client

import (
	"encoding/json"
	"keesvv/go-tcp-chat/internals/events"
	"keesvv/go-tcp-chat/internals/user"
	"net"
)

func Connect(opts Options) (*Connection, error) {
	ips, lookupErr := net.LookupIP(opts.Hostname)
	if lookupErr != nil {
		return &Connection{}, lookupErr
	}

	ip := ips[0]

	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   ip,
		Port: opts.Port,
	})

	if err != nil {
		return &Connection{}, err
	}

	user := &user.User{
		Nickname: opts.Nickname,
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
