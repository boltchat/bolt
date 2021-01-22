package tui

import "github.com/gdamore/tcell/v2"

const promptOffset int = 1

func displayPrompt(s tcell.Screen, input []rune) {
	w, h := s.Size()
	style := tcell.StyleDefault.Foreground(tcell.ColorYellow).Bold(true)
	y := h - promptOffset

	// Clear prompt line
	for i := 0; i < w; i++ {
		s.SetContent(i, y, ' ', nil, tcell.StyleDefault)
	}

	// Print prompt arrow
	s.SetContent(0, y, '>', nil, style)

	// Print user input
	for i := 0; i < len(input); i++ {
		s.SetContent(i+2, y, input[i], nil, tcell.StyleDefault)
	}

	// Draw a vertical line after input
	s.SetContent(len(input)+2, y, tcell.RuneVLine, nil, tcell.StyleDefault)

	s.Show()
}
