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
