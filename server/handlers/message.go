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
	"encoding/json"
	"strings"

	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/server/plugins"
	"github.com/bolt-chat/server/pools"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

func HandleMessage(p *pools.ConnPool, c *pools.Connection, e *events.BaseEvent) {
	msgEvt := &events.MessageEvent{}
	json.Unmarshal(*e.Raw, msgEvt)
	hexDecodeErr := plugins.GetManager().HookMessage(msgEvt, c)

	if hexDecodeErr != nil {
		c.Send(*events.NewErrorEvent(hexDecodeErr.Error()))
		return
	}

	sigReader := strings.NewReader(msgEvt.Message.Signature)
	sigDecoded, decodeErr := armor.Decode(sigReader)

	if decodeErr != nil {
		c.Send(*events.NewErrorEvent(decodeErr.Error())) // TODO
		return
	}

	pack, packErr := packet.Read(sigDecoded.Body)

	if packErr != nil {
		c.Send(*events.NewErrorEvent(packErr.Error())) // TODO
		return
	}

	_, ok := pack.(*packet.Signature)
	if !ok {
		c.Send(*events.NewErrorEvent("invalid_signature")) // TODO
		return
	}

	// pubKeyBin, hexDecodeErr := hex.DecodeString(fprint)
	// if hexDecodeErr != nil {
	// 	c.Send(*events.NewErrorEvent(hexDecodeErr.Error())) // TODO
	// 	return
	// }

	// pubKeyReader := bytes.NewReader(pubKeyBin)

	// pubKeyPack, pubKeyPackErr := packet.Read(pubKeyReader)
	// if pubKeyPackErr != nil {
	// 	c.Send(*events.NewErrorEvent(pubKeyPackErr.Error())) // TODO
	// 	return
	// }

	// pubKey, ok := pubKeyPack.(*packet.PublicKey)
	// if !ok {
	// 	c.Send(*events.NewErrorEvent("invalid_pubkey")) // TODO
	// 	return
	// }

	// pubKey.VerifySignature()

	p.Broadcast(msgEvt)
}
