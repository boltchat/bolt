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

package handlers

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/bolt-chat/protocol/errs"
	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/server/logging"
	"github.com/bolt-chat/server/plugins"
	"github.com/bolt-chat/server/pools"
)

/*
HandleConnection handles a TCP connection
during its entire lifespan.
*/
func HandleConnection(pool *pools.ConnPool, conn *pools.Connection) {
	for {
		// a := server.Listener{}
		b := make([]byte, 4096)

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

		if err != nil {
			conn.Send(*events.NewErrorEvent(errs.InvalidFormat))
			continue
		}

		if !conn.IsIdentified() && evt.Event.Type != events.JoinType {
			conn.Send(*events.NewErrorEvent(errs.Unidentified))
			continue
		}

		switch evt.Event.Type {
		case events.MessageType:
			msgEvt := &events.MessageEvent{}
			json.Unmarshal(b, msgEvt)
			err := plugins.GetManager().HookMessage(msgEvt, conn)

			if err != nil {
				conn.Send(*events.NewErrorEvent(err.Error()))
				continue
			}

			pool.Broadcast(msgEvt) // TODO: mutate and write
		case events.JoinType:
			joinEvt := &events.JoinEvent{}
			json.Unmarshal(b, joinEvt)
			conn.User = joinEvt.User

			motd, hasMotd := os.LookupEnv("MOTD") // Get MOTD env
			if hasMotd == true {
				// Send MOTD if env var is declared
				conn.Send(*events.NewMotdEvent(motd))
			}

			pool.Broadcast(joinEvt)
		default:
			conn.Send(*events.NewErrorEvent(errs.InvalidEvent))
		}
	}
}
