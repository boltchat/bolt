package events

import "time"

// Type represents an event type identifier.
type Type int

const (
	// MessageType TODO
	MessageType Type = iota
)

// Event represents a chat event.
type Event struct {
	Type      Type  `json:"t"`
	CreatedAt int64 `json:"c"`
}

// BaseEvent TODO
type BaseEvent struct {
	Event *Event `json:"e"`
}

func newEvent(t Type) *Event {
	return &Event{
		Type:      t,
		CreatedAt: time.Now().Unix(),
	}
}
