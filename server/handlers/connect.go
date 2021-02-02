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

package handlers

import (
	"bytes"
	"encoding/json"

	"github.com/bolt-chat/protocol/errs"
	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/server/logging"
	"github.com/bolt-chat/server/pools"
)

/*
HandleConnection handles a TCP connection
during its entire lifespan.
*/
func HandleConnection(pool *pools.ConnPool, conn *pools.Connection) {
	for {
		// Allocate 64KB for the event
		// TODO: automatically resize
		b := make([]byte, 65536)

		// Wait for and receive incoming events
		_, connErr := conn.Conn.Read(b)

		if connErr != nil {
			// Broadcast a disconnect message
			evt := *events.NewLeaveEvent(conn.User) // TODO:
			evtRaw, _ := json.Marshal(evt)
			pool.Broadcast(evt)
			logging.LogEvent(string(evtRaw))
			pool.RemoveFromPool(conn)
			return
		}

		// Trim empty bytes at the end
		b = bytes.TrimRight(b, "\x00")

		// Log raw events in debug mode
		logging.LogEvent(string(b))

		evt := &events.BaseEvent{}

		// Decode the event
		err := json.Unmarshal(b, evt)

		// Set raw byte value
		evt.Raw = &b

		if err != nil {
			conn.Send(*events.NewErrorEvent(errs.InvalidFormat))
			continue
		}

		if !conn.IsIdentified() && evt.Event.Type != events.JoinType {
			conn.Send(*events.NewErrorEvent(errs.Unidentified))
			continue
		}

		// Get and execute the corresponding event handler
		evtHandler := GetHandler(evt.Event.Type)
		evtHandler(pool, conn, evt)
	}
}
