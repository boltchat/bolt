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

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bolt-chat/server"
	"github.com/bolt-chat/server/plugins"
	"github.com/bolt-chat/server/util"
)

func main() {
	mgr := &plugins.PluginManager{}

	// Install plugins
	mgr.Install(
		plugins.RateLimiterPlugin{
			Amount: 5,
			Time:   time.Second,
		},
	)

	// Set the plugin manager
	plugins.SetManager(mgr)

	// Print ASCII banner
	util.PrintBanner()

	ipv4Bind := os.Getenv("BIND_IPV4")
	ipv6Bind := os.Getenv("BIND_IPV6")

	if ipv4Bind == "" {
		// Set default IPv4 bind to loopback address
		ipv4Bind = "127.0.0.1"
	}

	if ipv6Bind == "" {
		// Set default IPv6 bind to loopback address
		ipv6Bind = "::1"
	}

	listener := server.Listener{
		Bind: []server.Bind{
			{Address: ipv4Bind, Proto: "tcp4"},
			{Address: ipv6Bind, Proto: "tcp6"},
		},
		Port: 3300,
	}

	err := listener.Listen()
	if err != nil {
		panic(err)
	}

	// Exit on return
	fmt.Scanln()
}
