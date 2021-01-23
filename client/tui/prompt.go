package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/keesvv/bolt.chat/client/config"
)

type Mode int

const (
	MessageMode Mode = iota
)

var modes = map[Mode]string{
	MessageMode: "Msg",
}

func displayPrompt(s tcell.Screen, input []rune, mode Mode) {
	w, h := s.Size()
	style := tcell.StyleDefault.Foreground(tcell.ColorYellow).Bold(true)
	y := h - config.GetConfig().Prompt.HOffset

	modeStr := modes[mode]
	modeLen := len(modeStr)
	inputLen := len(input)

	arrowXPos := modeLen + 1
	inputXPos := inputLen + arrowXPos

	// Clear prompt line
	clearLine(s, y, w)

	// Print prompt mode
	for i := 0; i < modeLen; i++ {
		s.SetContent(i, y, rune(modeStr[i]), nil, tcell.StyleDefault.Bold(true))
	}

	// Print prompt arrow
	s.SetContent(arrowXPos, y, '>', nil, style)

	// Print user input
	for i := 0; i < inputLen; i++ {
		s.SetContent(i+arrowXPos+2, y, input[i], nil, tcell.StyleDefault)
	}

	// Draw cursor after input
	s.ShowCursor(inputXPos+2, y)

	s.Show()
}
