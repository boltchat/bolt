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

	"github.com/bolt-chat/protocol/errs"
	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/server/logging"
	"github.com/bolt-chat/server/pgp"
	"github.com/bolt-chat/server/plugins"
	"github.com/bolt-chat/server/pools"
)

func HandleMessage(p *pools.ConnPool, c *pools.Connection, e *events.BaseEvent) {
	msgEvt := &events.MessageEvent{}
	json.Unmarshal(*e.Raw, msgEvt)
	err := plugins.GetManager().HookMessage(msgEvt, c)

	if err != nil {
		c.SendError(err.Error())
		return
	}

	pubKey, verifyErr := pgp.VerifyMessageSignature(
		msgEvt.Message.Signature,
		c.User.PublicKey,
		msgEvt.Message.Content,
	)

	if verifyErr != nil {
		logging.LogDebug("Signature does not match.", nil)
		c.SendError(errs.SigVerifyFailed)
		return
	}

	logging.LogDebug("Signature matches.", nil)
	msgEvt.Message.Fingerprint = hex.EncodeToString(pubKey.Fingerprint[:])
	p.Broadcast(msgEvt)
}
