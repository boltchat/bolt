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
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/server/logging"
	"github.com/bolt-chat/server/plugins"
	"github.com/bolt-chat/server/pools"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

var publicKeyHex string = `// TODO`

func HandleMessage(p *pools.ConnPool, c *pools.Connection, e *events.BaseEvent) {
	msgEvt := &events.MessageEvent{}
	json.Unmarshal(*e.Raw, msgEvt)
	err := plugins.GetManager().HookMessage(msgEvt, c)

	if err != nil {
		c.Send(*events.NewErrorEvent(err.Error()))
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

	sig, ok := pack.(*packet.Signature)
	if !ok {
		c.Send(*events.NewErrorEvent("invalid_signature")) // TODO
		return
	}

	pubKeyRead := strings.NewReader(publicKeyHex)
	pubKeyDecode, pubKeyDecodeErr := armor.Decode(pubKeyRead)

	if pubKeyDecodeErr != nil {
		c.Send(*events.NewErrorEvent(pubKeyDecodeErr.Error())) // TODO
		return
	}

	pubKeyPack, pubKeyPackErr := packet.Read(pubKeyDecode.Body)

	if pubKeyPackErr != nil {
		c.Send(*events.NewErrorEvent(pubKeyPackErr.Error())) // TODO
		return
	}

	pubKey, ok := pubKeyPack.(*packet.PublicKey)
	if !ok {
		c.Send(*events.NewErrorEvent("invalid_pubkey")) // TODO
		return
	}

	hash := sig.Hash.New()
	_, hashErr := hash.Write([]byte(msgEvt.Message.Content))

	if hashErr != nil {
		c.Send(*events.NewErrorEvent(hashErr.Error())) // TODO
		return
	}

	verifyErr := pubKey.VerifySignature(hash, sig)

	if verifyErr != nil {
		logging.LogDebug("Signature does not match.", nil)
		c.Send(*events.NewErrorEvent("sig_verification_failed")) // TODO
		return
	}

	logging.LogDebug("Signature matches.", nil)
	msgEvt.Message.Fingerprint = hex.EncodeToString(pubKey.Fingerprint[:])
	p.Broadcast(msgEvt)
}
