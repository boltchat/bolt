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

package format

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bolt-chat/protocol/events"
	"github.com/fatih/color"
)

func FormatMessage(e *events.BaseEvent) string {
	msgEvt := &events.MessageEvent{}
	json.Unmarshal(*e.Raw, msgEvt)

	fprint := msgEvt.Message.Fingerprint

	// Use the last four characters of the fingerprint
	// for the user tag.
	tag := fprint[len(fprint)-4:]

	return fmt.Sprintf(
		"<%s#%s> %s",
		msgEvt.Message.User.Nickname,
		color.HiYellowString(strings.ToUpper(tag)),
		msgEvt.Message.Content,
	)
}
