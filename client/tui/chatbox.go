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
	"fmt"
	"strings"
	"time"

	"github.com/bolt-chat/client/config"
	"github.com/bolt-chat/client/format"
	"github.com/bolt-chat/protocol/events"
	"github.com/fatih/color"

	"github.com/gdamore/tcell/v2"
)

func printLine(s tcell.Screen, y int, str string) {
	/*
		I do not like this workaround at all, but at this
		point, I've given up on trying to find a better
		solution. Feel free to create a Pull Request if
		you're able to improve this.
	*/
	chars := []rune("\b\b" + str)

	s.SetContent(0, y, ' ', chars[1:], tcell.StyleDefault)
}

func printEvent(s tcell.Screen, y int, evt *events.BaseEvent) {
	var evtStr string
	var evtPrefix string
	var evtContent string

	// Convert event timestamp to `time.Time`
	timestamp := time.Unix(evt.Event.CreatedAt, 0)

	// Format the timestamp string
	timestampStr := strings.Join([]string{
		color.HiBlackString("["),
		timestamp.Format(time.Stamp),
		color.HiBlackString("]"),
	}, "")

	// Calculate timestamp length
	// TODO: refactor
	timestampLen := len(fmt.Sprintf(
		"[%s] ",
		timestamp.Format(time.Stamp),
	))

	if formatFunc, ok := format.FormatMap[evt.Event.Type]; ok {
		// Format the event
		evtContent = formatFunc(evt)

		// TODO: is this necessary?
		// /*
		//  Remove all control characters and non-printable
		//  characters from the event
		// */
		// evtContent = strings.TrimFunc(formattedEvt, func(r rune) bool {
		//  return unicode.IsControl(r) || !unicode.IsGraphic(r)
		// })
	} else {
		// No such formatter was found
		evtContent = fmt.Sprintf("unable to format event: %v", evt.Event.Type)
	}

	evtPrefix = timestampStr + " "
	evtStr = evtPrefix + evtContent

	for offset, line := range strings.Split(evtStr, "\n") {
		if offset > 0 {
			line = strings.Repeat(" ", timestampLen) + line
		}

		printLine(s, y+offset, line)
	}
}

func clearLine(s tcell.Screen, y int, w int) {
	// Clear every cell to `w`
	for x := 0; x < w; x++ {
		s.SetContent(x, y, ' ', nil, tcell.StyleDefault)
	}
}

func displayChatbox(s tcell.Screen, evtChannel chan *events.BaseEvent) {
	/*
		Preallocate a size of 50 for both the
		events slice and the buffer slice.
	*/
	evts := make([]*events.BaseEvent, 0, 50)
	buff := make([]*events.BaseEvent, 0, 50)

	for evt := range evtChannel {
		w, h := s.Size()
		hBuff := h - config.GetConfig().Prompt.HOffset

		// Append event to the events slice
		evts = append(evts, evt)

		if len(buff) < hBuff {
			// Append event to buffer
			buff = append(buff, evt)
		} else {
			// Remove first event from buffer and append
			buff = append(buff[1:], evt)
		}

		// Clear the buffer
		for y := 0; y < hBuff; y++ {
			clearLine(s, y, w)
		}

		// Append all events to the chatbox buffer
		for y, event := range buff {
			printEvent(s, y, event)
		}

		/*
			FIXME: This is for the time being. See issue #2
			for more information.
		*/
		s.Sync()
	}
}
