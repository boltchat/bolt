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
	"github.com/boltchat/client/errs"
	"github.com/boltchat/client/tui/chatbox"
	"github.com/boltchat/client/tui/prompt"
	"github.com/boltchat/lib/client"
	"github.com/boltchat/protocol/events"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

var screen tcell.Screen

/*
Display displays the TUI.
*/
func Display(c *client.Client, evts chan *events.Event) {
	// Register all encodings
	encoding.Register()

	// Channels
	clear := make(chan bool)
	termEvts := make(chan tcell.Event)

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
	go prompt.DisplayPrompt(s, c, termEvts, clear)
	go chatbox.DisplayChatbox(s, evts, clear)

	// Poll terminal events
	go func() {
		for {
			termEvts <- s.PollEvent()
		}
	}()
}

// Quit quits the TUI.
func Quit() {
	screen.Fini()
}
