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

package handlers

import (
	"bytes"
	"encoding/json"
	"net"
	"os"

	"github.com/bolt-chat/protocol"
	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/server/logging"
	"github.com/bolt-chat/util"
)

/*
HandleConnection handles a TCP connection
during its entire lifespan.
*/
func HandleConnection(pool *util.ConnPool, conn *net.TCPConn) {
	var user *protocol.User

	for {
		// a := server.Listener{}
		b := make([]byte, 4096)

		// Wait for and receive incoming events
		_, connErr := conn.Read(b)

		if connErr != nil {
			// Broadcast a disconnect message
			evt := *events.NewLeaveEvent(user) // TODO:
			evtRaw, _ := json.Marshal(evt)
			util.Broadcast(pool, evt) // TODO:
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
			util.WriteJson(conn, *events.NewErrorEvent("invalid_format"))
			conn.Close()
			return
		}

		switch evt.Event.Type {
		case events.MessageType:
			msgEvt := &events.MessageEvent{}
			json.Unmarshal(b, msgEvt)
			util.Broadcast(pool, msgEvt) // TODO: mutate and write
		case events.JoinType:
			joinEvt := &events.JoinEvent{}
			json.Unmarshal(b, joinEvt)
			user = joinEvt.User

			motd, hasMotd := os.LookupEnv("MOTD") // Get MOTD env
			if hasMotd == true {
				// Send MOTD if env var is declared
				util.WriteJson(conn, *events.NewMotdEvent(motd))
			}

			util.Broadcast(pool, joinEvt)
		default:
			util.WriteJson(conn, *events.NewErrorEvent("invalid_event"))
		}
	}
}
