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

	"github.com/boltchat/protocol"
	"github.com/boltchat/protocol/events"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
)

/*
SendMessage sends a message to an established
TCP connection.
*/
func (c *Client) SendMessage(m *protocol.Message) error {
	c.SendRaw(*events.NewMessageEvent(m))
	return nil
}

// SignMessage replaces the contents of a message with
// an Identity signature with the original contents embedded.
func (c *Client) SignMessage(m *protocol.Message) error {
	r := strings.NewReader(m.Content)
	buff := new(bytes.Buffer)

	err := openpgp.ArmoredDetachSignText(buff, c.Identity.Entity, r, &packet.Config{})
	if err != nil {
		return err
	}

	m.Signature = buff.String()
	return nil
}
