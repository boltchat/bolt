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

package client

import (
	"bytes"
	"strings"

	"github.com/bolt-chat/protocol"
	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/util"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
)

/*
SendMessage sends a message to an established
TCP connection.
*/
func (c *Client) SendMessage(m *protocol.Message) error {
	util.WriteJson(c.Conn, *events.NewMessageEvent(m))
	return nil
}

// SignMessage replaces the contents of a message with
// an Identity signature with the original contents embedded.
func (c *Client) SignMessage(m *protocol.Message) {
	r := strings.NewReader(m.Content)
	buff := new(bytes.Buffer)

	// TODO: do not write to stdout
	openpgp.ArmoredDetachSignText(buff, c.Identity.Entity, r, &packet.Config{})
	m.Content = buff.String()
}
