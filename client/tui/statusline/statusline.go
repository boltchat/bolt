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

package statusline

import (
	"github.com/boltchat/client"
	"github.com/boltchat/client/tui/util"
	"github.com/boltchat/protocol"
	"github.com/boltchat/protocol/events"
	"github.com/fatih/color"
	"github.com/gdamore/tcell/v2"
	"github.com/mitchellh/mapstructure"
)

var typingMap map[protocol.User]bool

func DisplayStatusLine(s tcell.Screen, evtsChan chan *events.Event) {
	typingMap = make(map[protocol.User]bool, 3)

	util.PrintLine(s, 0, 0, color.CyanString("Bolt Client v%s", client.Version.VersionString))
	s.Show()

	for evt := range evtsChan {
		if evt.Meta.Type != events.TypingType {
			continue
		}

		typingData := events.TypingData{}
		mapstructure.Decode(evt.Data, &typingData)

		if typingData.IsTyping {
			typingMap[*typingData.User] = true
		} else {
			delete(typingMap, *typingData.User)
		}

		w, _ := s.Size()
		util.ClearLine(s, 0, w)

		usernames := make([]string, 0, 3)
		for user := range typingMap {
			usernames = append(usernames, user.Nickname)
		}

		var typingText string

		if len(usernames) == 1 {
			typingText = color.CyanString("%s is typing...", usernames[0])
		} else if len(usernames) == 2 {
			typingText = color.CyanString("%s and %s are typing...", usernames[0], usernames[1])
		} else if len(usernames) > 2 {
			typingText = color.CyanString("Multiple people are typing...")
		}

		if typingText != "" {
			util.PrintLine(s, 0, 0, typingText)
		}

		s.Sync()
	}
}
