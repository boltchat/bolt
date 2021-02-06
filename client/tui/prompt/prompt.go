// Copyright 2021 The boltchat Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prompt

import (
	"os"
	"unicode"

	"github.com/boltchat/client/config"
	"github.com/boltchat/client/tui/util"
	"github.com/boltchat/lib/client"
	"github.com/gdamore/tcell/v2"
)

// FIXME: thread-unsafe
var input []rune

// FIXME: thread-unsafe
var mode Mode

func drawPrompt(s tcell.Screen) {
	w, h := s.Size()
	arrowStyle := tcell.StyleDefault.Foreground(tcell.ColorYellow).Bold(true)

	y := h - config.GetConfig().Prompt.HOffset

	modeStr := modeStrs[mode]
	modeLen := len(modeStr)
	inputLen := len(input)

	arrowXPos := modeLen + 1
	inputXPos := inputLen + arrowXPos

	// Clear prompt line
	util.ClearLine(s, y, w)

	// Print prompt mode
	for i := 0; i < modeLen; i++ {
		s.SetContent(i, y, rune(modeStr[i]), nil, tcell.StyleDefault.Bold(true))
	}

	// Print prompt arrow
	s.SetContent(arrowXPos, y, '>', nil, arrowStyle)

	// Print user input
	for i := 0; i < inputLen; i++ {
		s.SetContent(i+arrowXPos+2, y, input[i], nil, tcell.StyleDefault)
	}

	// Draw cursor after input
	s.ShowCursor(inputXPos+2, y)

	s.Show()
}

func handleEvents(s tcell.Screen, c *client.Client, termEvts chan tcell.Event, clear chan bool) {
	for termEvt := range termEvts {
		// Execute mode handlers
		modeHandlers[mode](s, c, termEvt)

		switch termEvt.(type) {
		case *tcell.EventKey:
			evt := termEvt.(*tcell.EventKey)

			if evt.Key() == tcell.KeyEscape ||
				evt.Key() == tcell.KeyCtrlC ||
				evt.Key() == tcell.KeyCtrlD {
				// Exit TUI
				s.Fini()
				os.Exit(0)
				return
			} else if evt.Key() == tcell.KeyCtrlL {
				go func() { clear <- true }()
			} else if evt.Key() == tcell.KeyBackspace2 {
				if len(input) > 0 {
					input = input[:len(input)-1]
				}
			} else if evt.Key() == tcell.KeyCtrlU {
				clearInput()
			} else if evt.Key() == tcell.KeyUp ||
				evt.Key() == tcell.KeyDown ||
				evt.Key() == tcell.KeyLeft ||
				evt.Key() == tcell.KeyRight ||
				evt.Key() == tcell.KeyHome ||
				evt.Key() == tcell.KeyEnd {
				// TODO: add logic
				break
			} else if evt.Rune() == '/' && len(input) == 0 {
				input = append(input, evt.Rune())
				mode = CommandMode
			} else {
				// Append the character if it's visible
				if unicode.IsGraphic(evt.Rune()) {
					input = append(input, evt.Rune())
				}
			}

			drawPrompt(s)
		}
	}
}

func DisplayPrompt(s tcell.Screen, c *client.Client, termEvts chan tcell.Event, clear chan bool) {
	// Initialize vars
	clearInput()
	mode = MessageMode

	// Draw initial (empty) prompt
	drawPrompt(s)

	// Handle terminal events
	go handleEvents(s, c, termEvts, clear)
}
