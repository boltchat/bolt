// bolt.chat
// Copyright (C) 2021  Kees van Voorthuizen
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

package util

import (
	"encoding/json"
	"net"

	"github.com/bolt-chat/server/logging"
)

type ConnPool []*net.TCPConn

func (c *ConnPool) AddToPool(conn *net.TCPConn) {
	// Append connection to pool
	*c = append(*c, conn)

	logging.LogDebug(
		"connection added to pool:",
		conn.RemoteAddr().String(),
	)

	logging.LogDebug("pool size:", len(*c))
}

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
			return
		}
	}
}

func WriteJson(conn *net.TCPConn, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	conn.Write(b)
}

func Broadcast(conns *ConnPool, data interface{}) {
	for _, conn := range *conns {
		WriteJson(conn, data)
	}
}
