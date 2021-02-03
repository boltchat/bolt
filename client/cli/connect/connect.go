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

package connect

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bolt-chat/client/cli/cmd"
	cliIdentity "github.com/bolt-chat/client/cli/identity"
	"github.com/bolt-chat/client/config"
	"github.com/bolt-chat/client/errs"
	"github.com/bolt-chat/client/identity"
	"github.com/bolt-chat/client/tui"
	"github.com/bolt-chat/lib/client"
	"github.com/bolt-chat/protocol/events"
	"github.com/fatih/color"
)

var ConnectCommand = &cmd.Command{
	Name:    "connect",
	Desc:    "Connects to a Bolt instance.",
	Usage:   "<host> [identity]",
	Handler: connectHandler,
}

type Args struct {
	Hostname string
	Port     int
	Identity string
}

// TODO: validation functions on commands
func parseArgs(args []string) (*Args, error) {
	// Set identity to 'default' by default
	identity := config.DefaultIdentity

	if len(args) < 1 {
		return nil, errors.New("no host given")
	}

	if len(args) > 1 {
		identity = args[1]
	}

	splitHost := strings.Split(args[0], ":")
	hostname := splitHost[0]

	// The default port
	port := 3300

	// Custom port number is specified
	if len(splitHost) == 2 {
		parsedPort, parseErr := strconv.ParseInt(splitHost[1], 10, 32)

		if parseErr != nil {
			return nil, errors.New("invalid host")
		}

		port = int(parsedPort)
	}

	return &Args{
		Hostname: hostname,
		Port:     port,
		Identity: identity,
	}, nil
}

func connectHandler(args []string) error {
	connectArgs, parseErr := parseArgs(args)
	if parseErr != nil {
		return parseErr
	}

	// Attempt to read the identity
	idEntry, identityErr := config.GetIdentityEntry(connectArgs.Identity)
	var id *identity.Identity

	// TODO: refactor
	if identityErr == config.ErrNoSuchIdentity {
		if !cliIdentity.AskCreate(connectArgs.Identity) {
			os.Exit(1)
		}

		var createErr error
		id, createErr = cliIdentity.CreateIdentity(connectArgs.Identity)

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
		Hostname: connectArgs.Hostname,
		Port:     connectArgs.Port,
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

	return nil
}
