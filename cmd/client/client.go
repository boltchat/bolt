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
	"fmt"
	"os"

	"github.com/bolt-chat/client/args"
	cliIdentity "github.com/bolt-chat/client/cli/identity"
	"github.com/bolt-chat/client/config"
	"github.com/bolt-chat/client/errs"
	"github.com/bolt-chat/client/identity"
	"github.com/bolt-chat/client/tui"
	"github.com/bolt-chat/lib/client"
	"github.com/bolt-chat/protocol/events"
	"github.com/fatih/color"
)

func main() {
	// Load the config
	config.LoadConfig()
	config.LoadIdentityList()

	args := args.GetArgs()

	// Attempt to read the identity
	idEntry, identityErr := config.GetIdentityEntry(args.Identity)
	var id *identity.Identity

	// TODO: refactor
	if identityErr == config.ErrNoSuchIdentity {
		if !cliIdentity.AskCreate(args.Identity) {
			os.Exit(1)
		}

		var createErr error
		id, createErr = cliIdentity.CreateIdentity(args.Identity)

		if createErr != nil {
			errs.Identity(createErr)
		}
	} else if identityErr != nil {
		errs.Identity(identityErr)
	} else {
		var loadErr error
		id, loadErr = identity.LoadIdentity(idEntry)

		if loadErr != nil {
			errs.Identity(loadErr)
		}
	}

	if id.Nickname == "" {
		errs.General(
			fmt.Sprintf(
				"It looks like you haven't set your nickname "+
					"yet.\nPlease do so by editing the %s field in %s.",
				color.HiYellowString("nickname"),
				config.IdentityFile.GetLocation(),
			),
		)
	}

	c := client.NewClient(client.Options{
		Hostname: args.Hostname,
		Port:     args.Port,
		Identity: id,
	})

	err := c.Connect()

	if err != nil {
		errs.Connect(err)
	}

	evts := make(chan *events.BaseEvent)

	serverClosed := make(chan bool)
	go c.ReadEvents(evts, serverClosed)
	go tui.Display(c, evts)

	// Quit when the server closes
	<-serverClosed
	tui.Quit()
	fmt.Println("The server closed.")
	os.Exit(0)
}
