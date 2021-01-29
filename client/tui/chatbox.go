// boltchat
// Copyright (C) 2021  The boltchat Authors
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

func printEvent(s tcell.Screen, w int, y int, evt *events.BaseEvent) int {
	// Convert event timestamp to `time.Time`
	timestamp := time.Unix(evt.Event.CreatedAt, 0)

	// Format the timestamp string
	timestampStr := strings.Join([]string{
		color.HiBlackString("["),
		timestamp.Format(time.Stamp),
		color.HiBlackString("]"),
	}, "")

	/*
		Calculate prefix length

		TODO: refactor. This is a temporary workaround
		because I have not yet found an optimal way of
		extracting control characters/ANSI colors from
		`timestampStr` and counting the length of that
		instead.
	*/
	prefixLen := len(fmt.Sprintf(
		"[%s] ",
		timestamp.Format(time.Stamp),
	))

	evtContent := format.Format(evt)
	evtPrefix := timestampStr + " "
	evtStr := evtPrefix + evtContent

	/*
		Preallocate one chunk because we're certain
		that there will always be at least one
		chunk in the `chunks` array.
	*/
	chunks := make([]string, 0, 1)

	// Split the event into an array of chunks
	for _, line := range strings.Split(evtStr, "\n") {
		chunks = append(chunks, splitChunks(line, w-prefixLen)...)
	}

	for offset, line := range chunks {
		if offset > 0 {
			line = strings.Repeat(" ", prefixLen) + line
		}

		printLine(s, y+offset, line)
	}

	return len(chunks) - 1
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
		yOffset := 0

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
			yOffset += printEvent(s, w, y+yOffset, event)
		}

		/*
			FIXME: This is for the time being. See issue #2
			for more information.
		*/
		s.Sync()
	}
}
