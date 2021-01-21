package client

import (
	"bufio"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/v2/encoding"
)

func displayPrompt(s tcell.Screen, input []rune) {
	w, h := s.Size()
	style := tcell.StyleDefault.Foreground(tcell.ColorYellow).Bold(true)
	y := h - 1

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

	s.Sync() // TODO: optimise
}

func displayChatbox(s tcell.Screen) {
	w, h := s.Size()

	for y := 0; y < h-1; y++ {
		for x := 0; x < w; x++ {
			s.SetContent(x, y, ' ', nil, tcell.StyleDefault)
		}
	}
}

/*
Prompt prompts the user for sending messages.
*/
func Prompt(r *bufio.Reader, conn *Connection) {
	encoding.Register()
	input := make([]rune, 0, 20)

	// Create a screen
	s, err := tcell.NewScreen()

	if err != nil {
		panic(err)
	}

	if err := s.Init(); err != nil {
		panic(err)
	}

	// Set style
	s.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite))
	displayPrompt(s, input)
	displayChatbox(s)

	for {
		switch ev := s.PollEvent().(type) {
		// case *tcell.EventResize:
		// 	s.Sync()
		// 	displayPrompt(s)
		// 	displayChatbox(s)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				// Exit prompt
				s.Fini()
				return
			} else if ev.Key() == tcell.KeyEnter {
				input = []rune{}
			} else if ev.Key() == tcell.KeyBackspace2 && len(input) > 0 {
				input = input[:len(input)-1]
			} else {
				input = append(input, ev.Rune())
			}

			displayPrompt(s, input)
		}
	}

	// msg := protocol.Message{
	// 	Content: content,
	// 	User:    &conn.User,
	// }

	// err := conn.SendMessage(&msg)
	// if err != nil {
	// 	panic(err)
	// }
}
