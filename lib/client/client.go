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
	Opts Options
}

func NewClient(opts Options) *Client {
	return &Client{
		User: protocol.User{
			Nickname: opts.Nickname,
		},
		Opts: opts,
	}
}

func (c *Client) Connect() error {
	ips, lookupErr := net.LookupIP(c.Opts.Hostname)
	if lookupErr != nil {
		return lookupErr
	}

	ip := ips[0]
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   ip,
		Port: c.Opts.Port,
	})

	if err != nil {
		return err
	}

	// Set the connection
	c.Conn = conn

	util.WriteJson(conn, *events.NewJoinEvent(&c.User))
	return nil
}
