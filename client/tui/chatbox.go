package tui

import (
	"github.com/bolt-chat/client/config"
	"github.com/bolt-chat/client/format"
	"github.com/bolt-chat/protocol/events"

	"github.com/gdamore/tcell/v2"
)

func printEvent(s tcell.Screen, y int, evt string) {
	/*
		I do not like this workaround at all, but at this
		point, I've given up on trying to find a better
		solution. Feel free to create a Pull Request if
		you're able to improve this.
	*/
	chars := []rune("\b\b" + evt)

	s.SetContent(0, y, ' ', chars[1:], tcell.StyleDefault)
}

func clearLine(s tcell.Screen, y int, w int) {
	// Clear every cell to `w`
	for x := 0; x < w; x++ {
		s.SetContent(x, y, ' ', nil, tcell.StyleDefault)
	}
}

func displayChatbox(s tcell.Screen, evtChannel chan *events.BaseEvent) {
	/*
		Preallocate a size of 50 for both the
		events slice and the buffer slice.
	*/
	evts := make([]*events.BaseEvent, 0, 50)
	buff := make([]*events.BaseEvent, 0, 50)

	for evt := range evtChannel {
		w, h := s.Size()
		hBuff := h - config.GetConfig().Prompt.HOffset

		// Append event to the events slice
		evts = append(evts, evt)

		if len(buff) < hBuff {
			// Append event to buffer
			buff = append(buff, evt)
		} else {
			// Remove first event from buffer and append
			buff = append(buff[1:], evt)
		}

		// Clear the buffer
		for y := 0; y < hBuff; y++ {
			clearLine(s, y, w)
		}

		// Append all events to the chatbox buffer
		for y, event := range buff {
			if formatFunc, ok := format.FormatMap[event.Event.Type]; ok {
				printEvent(s, y, formatFunc(event))
			}
		}

		/*
			FIXME: This is for the time being. See issue #2
			for more information.
		*/
		s.Sync()
	}
}
