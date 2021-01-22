package main

import (
	"github.com/keesvv/bolt.chat/client"
	"github.com/keesvv/bolt.chat/client/tui"
	"github.com/keesvv/bolt.chat/protocol/events"
)

func main() {
	conn, err := client.Connect(client.Options{
		Hostname: "localhost",
		Port:     3300,
		Nickname: "keesvv",
	})

	if err != nil {
		panic(err)
	}

	evts := make(chan *events.BaseEvent)

	go conn.ReadEvents(evts)
	tui.Display(conn, evts)
}
