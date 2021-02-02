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

package main

import (
	"os"
	"os/signal"
	"syscall"
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

	// Exit on syscall
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigs
}
