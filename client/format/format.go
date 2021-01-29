// boltchat
// Copyright (C) 2021  The boltchat Authors
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
