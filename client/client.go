package client

import (
	"net"

	"github.com/keesvv/bolt.chat/protocol"
	"github.com/keesvv/bolt.chat/protocol/events"
	"github.com/keesvv/bolt.chat/util"
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

	user := &protocol.User{
		Nickname: opts.Nickname,
	}

	util.WriteJson(conn, *events.NewJoinEvent(user))

	return &Connection{
		TCPConn: conn,
		User:    *user,
	}, nil
}
