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

package events

import "time"

// Type represents an event type identifier.
type Type string

// Event represents a general protocol event.
type Event struct {
	// The event identifier/type.
	Type Type `json:"t"`
	// The event creation date (client-side, untrustworthy).
	CreatedAt int64 `json:"c"`
	// The event receipt date (server-side, trustworthy).
	RecvAt int64 `json:"r,omitempty"` // TODO:
}

// BaseEvent represents a bare event without
// additional provided data.
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
