package main

import (
	"keesvv/go-tcp-chat/client"
	"keesvv/go-tcp-chat/client/tui"
	"keesvv/go-tcp-chat/protocol/events"
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
