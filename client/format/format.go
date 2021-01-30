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
	"fmt"

	"github.com/bolt-chat/protocol/events"
)

type formatHandler = func(e *events.BaseEvent) string

var formatMap = map[events.Type]formatHandler{
	events.MotdType:    FormatMotd,
	events.MessageType: FormatMessage,
	events.ErrorType:   FormatError,
	events.JoinType:    FormatJoin,
	events.LeaveType:   FormatLeave,
}

// Format formats an event in a human-readable format.
func Format(evt *events.BaseEvent) string {
	if formatFunc, ok := formatMap[evt.Event.Type]; ok {
		return formatFunc(evt)
	}
	return fmt.Sprintf("unable to format event: %v", evt.Event.Type)
}
