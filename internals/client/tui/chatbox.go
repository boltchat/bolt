package tui

import (
	"github.com/gdamore/tcell"
)

const promptOffset int = 1

func displayChatbox(s tcell.Screen, evts chan string) {
	w, h := s.Size()
	events := make([]string, 0, 50)

	for evt := range evts {
		events = append(events, evt)

		// Clear the chatbox
		for y := 0; y < (h - promptOffset); y++ {
			for x := 0; x < w; x++ {
				s.SetContent(x, y, ' ', nil, tcell.StyleDefault)
			}
		}

		// Append all event to the chatbox
		for y, line := range events {
			chars := []rune(line)

			for x, ch := range chars {
				s.SetContent(x, y, ch, nil, tcell.StyleDefault)
			}
		}

		s.Show()
	}
}
