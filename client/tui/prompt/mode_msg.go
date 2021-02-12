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

package prompt

import (
	"strings"
	"time"
	"unicode"

	"github.com/boltchat/client/errs"
	"github.com/boltchat/lib/client"
	"github.com/boltchat/protocol"
	"github.com/gdamore/tcell/v2"
)

const typingDuration = time.Second * 2

var typingTimer *time.Timer

func handleMessageMode(s tcell.Screen, c *client.Client, evt tcell.Event) {
	key, ok := evt.(*tcell.EventKey)
	if !ok {
		return
	}

	// TODO: clean up this mess
	// too tired for this now
	if key.Key() == tcell.KeyEnter {
		if typingTimer != nil {
			typingTimer.Stop()
			typingTimer = nil
		}

		body := strings.TrimSpace(string(input))

		if len(body) < 1 {
			return
		}

		// Disable typing indicator before sending message
		c.SetTyping(false)
		sendMessage(body, c)
	} else if unicode.IsGraphic(key.Rune()) {
		if typingTimer == nil {
			// User starts typing
			c.SetTyping(true)
			typingTimer = time.AfterFunc(typingDuration, func() {
				c.SetTyping(false)
				typingTimer = nil
			})
		} else {
			// User is still typing
			typingTimer.Reset(typingDuration)
		}
	}
}

func sendMessage(body string, c *client.Client) {
	msg := protocol.Message{
		Content: body,
		User: &protocol.User{
			Nickname: c.Identity.Nickname, // TODO
		},
	}

	signErr := c.SignMessage(&msg)
	if signErr != nil {
		errs.Emerg(signErr)
	}

	sendErr := c.SendMessage(&msg)
	if sendErr != nil {
		errs.Emerg(sendErr)
	}

	// Clear input
	clearInput()
}
