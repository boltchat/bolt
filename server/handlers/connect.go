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
	"github.com/bolt-chat/protocol/errs"
	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/server/pools"
)

/*
HandleConnection handles a TCP connection
during its entire lifespan.
*/
func HandleConnection(pool *pools.ConnPool, conn *pools.Connection) {
	for {
		evt := &events.BaseEvent{}

		// Wait for and receive incoming events
		connErr := conn.Read(evt)

		if connErr != nil {
			// Broadcast a disconnect message
			pool.Broadcast(*events.NewLeaveEvent(conn.User))
			pool.RemoveFromPool(conn)
			return
		}

		// TODO:
		// if err != nil {
		// 	conn.SendError(errs.InvalidFormat)
		// 	continue
		// }

		if !conn.IsIdentified() && evt.Meta.Type != events.JoinType {
			conn.SendError(errs.Unidentified)
			continue
		}

		// Get and execute the corresponding event handler
		evtHandler := GetHandler(evt.Meta.Type)
		evtHandler(pool, conn, evt)
	}
}
