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

	"github.com/bolt-chat/protocol/events"
	"github.com/fatih/color"
	"github.com/gdamore/tcell/v2"
)

func FormatLeave(e *events.BaseEvent) string {
	leaveEvt := &events.JoinEvent{}
	json.Unmarshal(*e.Raw, leaveEvt)

	return color.HiMagentaString(
		fmt.Sprintf("%s %s left the room.", string(tcell.RuneDiamond), leaveEvt.User.Nickname),
	)
}
