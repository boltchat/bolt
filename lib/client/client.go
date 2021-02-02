// Copyright 2021 The boltchat Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"net"

	"github.com/bolt-chat/client/identity"
	"github.com/bolt-chat/protocol"
	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/util"
)

type Client struct {
	Conn     *net.TCPConn // TODO: make private
	Identity *identity.Identity
	Opts     Options
}

func NewClient(opts Options) *Client {
	return &Client{
		Identity: opts.Identity,
		Opts:     opts,
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

	util.WriteJson(conn, *events.NewJoinEvent(&protocol.User{
		Nickname: c.Identity.Nickname, // TODO:
	}))
	return nil
}
