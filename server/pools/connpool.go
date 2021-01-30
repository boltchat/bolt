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

package pools

import (
	"github.com/bolt-chat/server/logging"
)

// ConnPool represents a group of connections.
type ConnPool []*Connection

func (c *ConnPool) logPoolSize() {
	logging.LogDebug("pool size:", len(*c))
}

// AddToPool adds a new connection to the current pool.
func (c *ConnPool) AddToPool(conn *Connection) {
	*c = append(*c, conn)

	logging.LogDebug(
		"connection added to pool:",
		conn.Conn.RemoteAddr().String(),
	)

	c.logPoolSize()
}

// RemoveFromPool removes an existing connection
// from the current pool.
func (c *ConnPool) RemoveFromPool(conn *Connection) {
	// Range through pool
	for i, curConn := range *c {
		// Target connection is found
		if curConn == conn {
			/*
				This removes the connection from the pool by its
				respective index.

				`i` represents the index of the matched connection.
				The first arg represents a slice of the pool that
				ends at the index of the connection in question.
				The second arg represents a slice of the pool that
				starts from the index of the connection in question + 1.
			*/
			*c = append((*c)[:i], (*c)[i+1:]...)
			break
		}
	}

	logging.LogDebug(
		"connection removed from pool:",
		conn.Conn.RemoteAddr().String(),
	)

	c.logPoolSize()
}

// Broadcast emits data to all connections that are
// present in the pool.
func (c *ConnPool) Broadcast(data interface{}) {
	for _, conn := range *c {
		conn.Send(data)
	}
}
