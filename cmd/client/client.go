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
	"github.com/bolt-chat/client/config"
	"github.com/bolt-chat/client/errs"
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

	identity, identityErr := config.GetIdentity(args.Identity)
	if identityErr != nil {
		errs.Emerg(identityErr)
	}

	if identity.Nickname == "" {
		errs.General(
			fmt.Sprintf(
				"It looks like you haven't set your nickname "+
					"yet.\nPlease do so by editing the %s field in %s.",
				color.HiYellowString("nickname"),
				config.GetIdentityLocation(),
			),
		)
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

	serverClosed := make(chan bool)
	go c.ReadEvents(evts, serverClosed)
	go tui.Display(c, evts)

	// Quit when the server closes
	<-serverClosed
	tui.Quit()
	fmt.Println("The server closed.")
	os.Exit(0)
}
