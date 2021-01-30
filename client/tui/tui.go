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

package tui

import (
	"os"
	"strings"

	"github.com/bolt-chat/client/errs"
	"github.com/bolt-chat/lib/client"
	"github.com/bolt-chat/protocol"
	"github.com/bolt-chat/protocol/events"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

var screen tcell.Screen

/*
Display displays the TUI.
*/
func Display(c *client.Client, evts chan *events.BaseEvent) {
	encoding.Register()
	input := make([]rune, 0, 20)
	mode := MessageMode

	// Create a screen
	s, err := tcell.NewScreen()
	screen = s

	if err != nil {
		errs.Emerg(err)
	}

	// Initialize the screen
	if err := s.Init(); err != nil {
		errs.Emerg(err)
	}

	// Set default style
	s.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite))

	// Display prompt and chatbox
	displayPrompt(s, input, mode)
	go displayChatbox(s, evts)

	for {
		switch ev := s.PollEvent().(type) {
		// case *tcell.EventResize:
		// 	s.Sync()
		// 	displayPrompt(s)
		// 	displayChatbox(s)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape ||
				ev.Key() == tcell.KeyCtrlC ||
				ev.Key() == tcell.KeyCtrlD {
				// Exit TUI
				Quit()
				os.Exit(0)
				return
			} else if ev.Key() == tcell.KeyEnter {
				if len(strings.TrimSpace(string(input))) < 1 {
					break
				}

				msg := protocol.Message{
					Content: string(input),
					User:    &c.User,
				}

				err := c.SendMessage(&msg)
				if err != nil {
					errs.Emerg(err)
				}

				input = []rune{}
			} else if ev.Key() == tcell.KeyBackspace2 {
				if len(input) > 0 {
					input = input[:len(input)-1]
				}
			} else if ev.Key() == tcell.KeyCtrlU {
				input = []rune{}
			} else if ev.Key() == tcell.KeyUp ||
				ev.Key() == tcell.KeyDown ||
				ev.Key() == tcell.KeyLeft ||
				ev.Key() == tcell.KeyRight ||
				ev.Key() == tcell.KeyHome ||
				ev.Key() == tcell.KeyEnd {
				// TODO: add logic
				break
			} else {
				input = append(input, ev.Rune())
			}

			displayPrompt(s, input, mode)
		}
	}
}

func Quit() {
	screen.Fini()
}
