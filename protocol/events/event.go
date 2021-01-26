// bolt.chat
// Copyright (C) 2021  Kees van Voorthuizen
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

package events

import "time"

// Type represents an event type identifier.
type Type string

const (
	// MessageType is the event type used for messages.
	MessageType Type = "msg"
	// JoinType is the event type used for new connections.
	JoinType Type = "join"
	// LeaveType is the event type used for disconnects.
	LeaveType Type = "leave"
	// ErrorType is the event type used for error reporting.
	ErrorType Type = "err"
	// MotdType is the event type used for the Message-of-the-Day (MOTD).
	MotdType Type = "motd"
)

// Event represents a general protocol event.
type Event struct {
	// The event identifier/type.
	Type Type `json:"t"`
	// The event creation date (client-side, untrustworthy).
	CreatedAt int64 `json:"c"`
	// The event receipt date (server-side, trustworthy).
	RecvAt int64 `json:"r,omitempty"` // TODO:
}

// BaseEvent TODO
type BaseEvent struct {
	Event *Event  `json:"e"`
	Raw   *[]byte `json:"-"`
}

// NewBaseEvent TODO
func NewBaseEvent(t Type) *BaseEvent {
	return &BaseEvent{
		Event: &Event{
			Type:      t,
			CreatedAt: time.Now().Unix(),
		},
	}
}
