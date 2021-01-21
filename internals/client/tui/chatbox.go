package tui

import "github.com/gdamore/tcell"

const promptOffset int = 1

func displayChatbox(s tcell.Screen) {
	w, h := s.Size()

	for y := 0; y < (h - promptOffset); y++ {
		for x := 0; x < w; x++ {
			s.SetContent(x, y, ' ', nil, tcell.StyleDefault)
		}
	}
}
