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
	"encoding/json"
	"net"

	"github.com/bolt-chat/protocol"
	"github.com/bolt-chat/server/logging"
)

// Connection TODO
type Connection struct {
	Conn    *net.TCPConn
	User    *protocol.User
	encoder *json.Encoder
	decoder *json.Decoder
}

// NewConnection TODO
func NewConnection(conn *net.TCPConn, user *protocol.User) *Connection {
	enc := json.NewEncoder(conn)
	dec := json.NewDecoder(conn)

	dec.DisallowUnknownFields()

	return &Connection{
		Conn:    conn,
		User:    user,
		encoder: enc,
		decoder: dec,
	}
}

// Send TODO
func (c *Connection) Send(data interface{}) {
	err := c.encoder.Encode(data)

	if err != nil {
		logging.LogError(err.Error())
	}
}

// Read TODO
func (c *Connection) Read(data interface{}) {
	err := c.decoder.Decode(data)

	if err != nil {
		logging.LogError(err.Error())
	}
}

// IsIdentified TODO
func (c *Connection) IsIdentified() bool {
	return c.User != nil
}
