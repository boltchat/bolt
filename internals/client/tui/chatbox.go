package tui

import (
	"keesvv/go-tcp-chat/internals/client/format"
	"keesvv/go-tcp-chat/internals/protocol/events"

	"github.com/gdamore/tcell"
)

const promptOffset int = 1

func printLine(s tcell.Screen, y int, line string) {
	chars := []rune(line)

	for x, ch := range chars {
		s.SetContent(x, y, ch, nil, tcell.StyleDefault)
	}
}

func displayChatbox(s tcell.Screen, evtChannel chan *events.BaseEvent) {
	w, h := s.Size()
	evts := make([]*events.BaseEvent, 0, 50)

	for evt := range evtChannel {
		evts = append(evts, evt)

		// Clear the chatbox
		for y := 0; y < (h - promptOffset); y++ { // TODO
			for x := 0; x < w; x++ {
				s.SetContent(x, y, ' ', nil, tcell.StyleDefault)
			}
		}

		// Append all event to the chatbox
		for y, event := range evts {
			formatMap := map[events.Type]format.FormatHandler{
				events.MotdType:    func(e *events.BaseEvent) string { return "TODO" },
				events.MessageType: format.FormatMessage,
			}

			printLine(s, y, formatMap[event.Event.Type](event))
		}

		s.Show()
	}
}
