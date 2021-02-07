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

	"github.com/boltchat/lib/client"
	"github.com/boltchat/protocol/events"
	"github.com/gdamore/tcell/v2"
)

func handleCommandMode(s tcell.Screen, c *client.Client, evt tcell.Event) {
	key, ok := evt.(*tcell.EventKey)
	if !ok {
		return
	}

	if (key.Key() == tcell.KeyBackspace || key.Key() == tcell.KeyBackspace2) &&
		len(input) == 1 &&
		input[0] == '/' {
		// Return to message mode
		mode = MessageMode
		return
	}

	if key.Key() == tcell.KeyEnter && len(input) > 1 {
		sendCommand(c)
		mode = MessageMode
		clearInput()
	}
}

func sendCommand(c *client.Client) {
	cmdSplit := strings.Split(string(input[1:]), " ")
	cmd := cmdSplit[0]
	args := cmdSplit[1:]

	c.SendCommand(&events.CommandData{
		Command: cmd,
		Args:    args,
	})
}
