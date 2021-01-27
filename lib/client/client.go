// bolt.chat
// Copyright (C) 2021  The bolt.chat Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package client

import (
	"net"

	"github.com/bolt-chat/protocol"
	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/util"
)

type Client struct {
	Conn *net.TCPConn // TODO: make private
	User protocol.User
}

// TODO: rename to `NewClient`, add separate `Connect` method
func Connect(opts Options /*TODO*/) (*Client, error) {
	ips, lookupErr := net.LookupIP(opts.Hostname)
	if lookupErr != nil {
		return &Client{}, lookupErr
	}

	ip := ips[0]

	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   ip,
		Port: opts.Port,
	})

	if err != nil {
		return &Client{}, err
	}

	user := &protocol.User{
		Nickname: opts.Nickname,
	}

	util.WriteJson(conn, *events.NewJoinEvent(user))

	return &Client{
		Conn: conn,
		User: *user,
	}, nil
}
