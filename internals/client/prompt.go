package client

import (
	"bufio"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/v2/encoding"
)

func displayPrompt(s tcell.Screen) {
	_, h := s.Size()
	style := tcell.StyleDefault.Foreground(tcell.ColorYellow).Bold(true)
	s.SetContent(0, h-1, '>', nil, style)
	s.Show()
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
	displayPrompt(s)
	displayChatbox(s)

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
			displayPrompt(s)
			displayChatbox(s)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				// Exit prompt
				s.Fini()
				return
			}
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
