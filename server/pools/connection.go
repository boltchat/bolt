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
	"encoding/json"
	"net"

	"github.com/bolt-chat/protocol"
	"github.com/bolt-chat/protocol/events"
)

// Connection TODO
type Connection struct {
	Conn    *net.TCPConn
	User    *protocol.User
	Data    map[string]interface{}
	encoder *json.Encoder
	decoder *json.Decoder
}

// NewConnection TODO
func NewConnection(conn *net.TCPConn, user *protocol.User) *Connection {
	enc := json.NewEncoder(conn)
	dec := json.NewDecoder(conn)

	return &Connection{
		Conn:    conn,
		User:    user,
		Data:    make(map[string]interface{}, 0),
		encoder: enc,
		decoder: dec,
	}
}

// Send TODO
func (c *Connection) Send(data interface{}) error {
	err := c.encoder.Encode(data)
	return err
}

// SendError TODO
func (c *Connection) SendError(err string) error {
	return c.Send(*events.NewErrorEvent(err))
}

// Receive TODO
func (c *Connection) Receive(data interface{}) error {
	err := c.decoder.Decode(data)
	return err
}

// Close closes the connection.
func (c *Connection) Close() error {
	return c.Conn.Close()
}

// IsIdentified TODO
func (c *Connection) IsIdentified() bool {
	return c.User != nil
}
