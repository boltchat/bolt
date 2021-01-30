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

package client

import (
	"bytes"
	"encoding/json"

	"github.com/bolt-chat/client/errs"
	"github.com/bolt-chat/protocol/events"
)

func (c *Client) ReadEvents(evts chan *events.BaseEvent, closed chan bool) {
	for {
		// Allocate 64KB for the event
		// TODO: automatically resize
		b := make([]byte, 65536)
		_, err := c.Conn.Read(b)

		if err != nil {
			closed <- true
			return
		}

		b = bytes.TrimRight(b, "\x00")

		evt := &events.BaseEvent{}
		jsonErr := json.Unmarshal(b, evt)

		evt.Raw = &b

		if jsonErr != nil {
			errs.Emerg(err)
		}

		evts <- evt
	}
}
