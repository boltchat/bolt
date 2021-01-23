package main

import (
	"github.com/keesvv/bolt.chat/client"
	"github.com/keesvv/bolt.chat/client/config"
	"github.com/keesvv/bolt.chat/client/tui"
	"github.com/keesvv/bolt.chat/protocol/events"
)

func main() {
	// Load the config
	config.LoadConfig()
	args := client.GetArgs()

	conn, err := client.Connect(client.Options{
		Hostname: args.Hostname,
		Port:     args.Port,
		Nickname: args.Nickname,
	})

	if err != nil {
		panic(err)
	}

	evts := make(chan *events.BaseEvent)

	go conn.ReadEvents(evts)
	tui.Display(conn, evts)
}
