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
	"github.com/bolt-chat/client/args"
	"github.com/bolt-chat/client/config"
	"github.com/bolt-chat/client/errs"
	"github.com/bolt-chat/client/tui"
	"github.com/bolt-chat/lib/client"
	"github.com/bolt-chat/protocol/events"
)

func main() {
	// Load the config
	config.LoadConfig()
	config.LoadIdentityList()

	args := args.GetArgs()

	identity, identityErr := config.GetIdentity(args.Identity)

	if identityErr != nil {
		errs.Emerg(identityErr)
	}

	c := client.NewClient(client.Options{
		Hostname: args.Hostname,
		Port:     args.Port,
		Nickname: identity.Nickname,
	})

	err := c.Connect()

	if err != nil {
		errs.Connect(err)
	}

	evts := make(chan *events.BaseEvent)

	go c.ReadEvents(evts)
	tui.Display(c, evts)
}
