package tui

import (
	"keesvv/go-tcp-chat/internals/client/format"
	"keesvv/go-tcp-chat/internals/protocol/events"

	"github.com/gdamore/tcell/v2"
)

const promptOffset int = 1

func printLine(s tcell.Screen, y int, line string) {
	/*
		I do not like this workaround at all, but at this
		point, I've given up on trying to find a better
		solution. Feel free to create a Pull Request if
		you're able to improve this.
	*/
	chars := []rune("\b\b" + line)

	s.SetContent(0, y, ' ', chars[1:], tcell.StyleDefault)
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
				events.MotdType:    format.FormatMotd,
				events.MessageType: format.FormatMessage,
				events.ErrorType:   format.FormatError,
			}

			printLine(s, y, formatMap[event.Event.Type](event))
		}

		s.Show()
	}
}
