// bolt.chat
// Copyright (C) 2021  The bolt.chat Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package tui

import (
	"github.com/bolt-chat/client/config"
	"github.com/gdamore/tcell/v2"
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
