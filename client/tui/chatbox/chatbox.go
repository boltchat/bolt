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

package chatbox

import (
	"fmt"
	"strings"
	"time"

	"github.com/boltchat/client/config"
	"github.com/boltchat/client/format"
	"github.com/boltchat/client/tui/util"
	"github.com/boltchat/protocol/events"
	"github.com/fatih/color"

	"github.com/gdamore/tcell/v2"
)

type EventBuffer = []*events.Event

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
		chunks = append(chunks, util.SplitChunks(line, w-prefixLen)...)
	}

	for offset, line := range chunks {
		if offset > 0 {
			line = strings.Repeat(" ", prefixLen) + line
		}

		util.PrintLine(s, 0, y+offset, line)
	}

	return len(chunks) - 1
}

func getBufferDimensions(s tcell.Screen) (h, yOffset int) {
	// Get heights & offsets
	_, sHeight := s.Size()
	statusHeight := config.GetConfig().StatusLine.Height
	promptHOffset := config.GetConfig().Prompt.HOffset

	h = sHeight - promptHOffset - statusHeight
	yOffset = statusHeight
	return
}

func clearBuffer(s tcell.Screen) {
	hBuff, yOffset := getBufferDimensions(s)
	w, _ := s.Size()

	// Clear the buffer
	for y := 0; y < hBuff; y++ {
		util.ClearLine(s, y+yOffset, w)
	}
}

func addToBuffer(evt *events.Event, hBuff int, buff EventBuffer) EventBuffer {
	if len(buff) > hBuff {
		// Remove first event from buffer and append
		return append(buff[1:], evt)
	}

	// Append event to buffer
	return append(buff, evt)
}

func DisplayChatbox(
	s tcell.Screen,
	evtChannel chan *events.Event,
	clear chan bool,
) {
	/*
		Preallocate a size of 50 for both the
		events slice and the buffer slice.
	*/
	// evts := make([]*events.Event, 0, 50)
	buff := make(EventBuffer, 0, 50)

	go func() {
		for evt := range evtChannel {
			w, _ := s.Size()
			hBuff, yOffset := getBufferDimensions(s)

			// evts = append(evts, evt)

			// Append event to the events slice if it has
			// a formatter.
			if format.HasFormat(evt) {
				buff = addToBuffer(evt, hBuff, buff)
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
