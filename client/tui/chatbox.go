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
	"fmt"
	"strings"
	"time"

	"github.com/bolt-chat/client/config"
	"github.com/bolt-chat/client/format"
	"github.com/bolt-chat/protocol/events"
	"github.com/fatih/color"

	"github.com/gdamore/tcell/v2"
)

func printEvent(s tcell.Screen, w int, y int, evt *events.Event) int {
	// Convert event timestamp to `time.Time`
	timestamp := time.Unix(evt.Meta.CreatedAt, 0)

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

func clearBuffer(s tcell.Screen) {
	w, h := s.Size()
	hBuff := h - config.GetConfig().Prompt.HOffset

	// Clear the buffer
	for y := 0; y < hBuff; y++ {
		clearLine(s, y, w)
	}
}

func displayChatbox(
	s tcell.Screen,
	evtChannel chan *events.Event,
	clear chan bool,
) {
	/*
		Preallocate a size of 50 for both the
		events slice and the buffer slice.
	*/
	evts := make([]*events.Event, 0, 50)
	buff := make([]*events.Event, 0, 50)

	go func() {
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
			clearBuffer(s)

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
	}()

	go func() {
		for c := range clear {
			if c {
				buff = make([]*events.Event, 0, 50)
				clearBuffer(s)
				s.Sync()
			}
		}
	}()
}
