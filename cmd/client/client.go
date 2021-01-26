// bolt.chat
// Copyright (C) 2021  Kees van Voorthuizen
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
	"github.com/bolt-chat/client"
	"github.com/bolt-chat/client/config"
	"github.com/bolt-chat/client/tui"
	"github.com/bolt-chat/protocol/events"
)

func main() {
	// Load the config
	config.LoadConfig()
	config.LoadIdentityList()

	args := client.GetArgs()

	identity, identityErr := config.GetIdentity(args.Identity)

	if identityErr != nil {
		panic(identityErr)
	}

	conn, err := client.Connect(client.Options{
		Hostname: args.Hostname,
		Port:     args.Port,
		Nickname: identity.Nickname,
	})

	if err != nil {
		panic(err)
	}

	evts := make(chan *events.BaseEvent)

	go conn.ReadEvents(evts)
	tui.Display(conn, evts)
}
