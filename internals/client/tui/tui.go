package tui

import (
	"keesvv/go-tcp-chat/internals/client"
	"keesvv/go-tcp-chat/internals/protocol"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

/*
Display displays the TUI.
*/
func Display(conn *client.Connection, evts chan string) {
	encoding.Register()
	input := make([]rune, 0, 20)

	// Create a screen
	s, err := tcell.NewScreen()

	if err != nil {
		panic(err)
	}

	// Initialize the screen
	if err := s.Init(); err != nil {
		panic(err)
	}

	// Set default style
	s.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite))

	// Display prompt and chatbox
	displayPrompt(s, input)
	go displayChatbox(s, evts)

	for {
		switch ev := s.PollEvent().(type) {
		// case *tcell.EventResize:
		// 	s.Sync()
		// 	displayPrompt(s)
		// 	displayChatbox(s)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				// Exit TUI
				s.Fini()
				return
			} else if ev.Key() == tcell.KeyEnter {
				if len(strings.TrimSpace(string(input))) < 1 {
					break
				}

				msg := protocol.Message{
					Content: string(input),
					User:    &conn.User,
				}

				err := conn.SendMessage(&msg)
				if err != nil {
					panic(err)
				}

				input = []rune{}
			} else if ev.Key() == tcell.KeyBackspace2 {
				if len(input) > 0 {
					input = input[:len(input)-1]
				}
			} else if ev.Key() == tcell.KeyCtrlU {
				input = []rune{}
			} else {
				input = append(input, ev.Rune())
			}

			displayPrompt(s, input)
		}
	}
}
