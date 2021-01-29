// boltchat
// Copyright (C) 2021  The boltchat Authors
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

// TODO: refactor
func (c *Client) Connect() error {
	var ip net.IP
	var port int = c.Opts.Port

	if parsedIP := net.ParseIP(c.Opts.Hostname); parsedIP != nil {
		ip = parsedIP
	}

	if ip == nil {
		_, srvAddrs, _ := net.LookupSRV("bolt", "tcp", c.Opts.Hostname)

		if len(srvAddrs) > 0 {
			targetIps, _ := net.LookupIP(srvAddrs[0].Target)
			ip = targetIps[0]
			port = int(srvAddrs[0].Port)
		}
	}

	if ip == nil {
		ips, lookupErr := net.LookupIP(c.Opts.Hostname)
		if lookupErr != nil {
			return lookupErr
		}

		ip = ips[0]
	}

	conn, dialErr := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   ip,
		Port: port,
	})

	if dialErr != nil {
		return dialErr
	}

	// Set the connection
	c.Conn = conn

	util.WriteJson(conn, *events.NewJoinEvent(&c.User))
	return nil
}
