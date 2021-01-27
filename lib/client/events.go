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

package client

import (
	"bytes"
	"encoding/json"

	"github.com/bolt-chat/protocol/events"
)

func (c *Client) ReadEvents(evts chan *events.BaseEvent) error {
	for {
		b := make([]byte, 4096)
		_, err := c.Conn.Read(b)

		if err != nil {
			return err
		}

		b = bytes.TrimRight(b, "\x00")
		evt := &events.BaseEvent{}
		jsonErr := json.Unmarshal(b, evt)

		evt.Raw = &b

		if jsonErr != nil {
			return err
		}

		go func() {
			evts <- evt
		}()
	}
}
