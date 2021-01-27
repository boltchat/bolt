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

package pools

import (
	"net"

	"github.com/bolt-chat/server/logging"
	"github.com/bolt-chat/util"
)

// ConnPool represents a group of connections.
type ConnPool []*net.TCPConn

func (c *ConnPool) logPoolSize() {
	logging.LogDebug("pool size:", len(*c))
}

// AddToPool adds a new connection to the current pool.
func (c *ConnPool) AddToPool(conn *net.TCPConn) {
	*c = append(*c, conn)

	logging.LogDebug(
		"connection added to pool:",
		conn.RemoteAddr().String(),
	)

	c.logPoolSize()
}

// RemoveFromPool removes an existing connection
// from the current pool.
func (c *ConnPool) RemoveFromPool(conn *net.TCPConn) {
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
		conn.RemoteAddr().String(),
	)

	c.logPoolSize()
}

// Broadcast emits data to all connections that are
// present in the pool.
func (c *ConnPool) Broadcast(data interface{}) {
	for _, conn := range *c {
		util.WriteJson(conn, data)
	}
}
