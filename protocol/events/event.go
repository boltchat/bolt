package events

import "time"

// Type represents an event type identifier.
type Type int

const (
	// MessageType TODO
	MessageType Type = iota
	// JoinType TODO
	JoinType Type = iota
	// LeaveType TODO
	LeaveType Type = iota
	// ErrorType TODO
	ErrorType Type = iota
	// MotdType TODO
	MotdType Type = iota
)

// Event represents a general protocol event.
type Event struct {
	// The event identifier/type.
	Type Type `json:"t"`
	// The event creation date (client-side, untrustworthy)
	CreatedAt int64 `json:"c"`
	// The event reception date (server-side, trustworthy)
	RecvAt int64 `json:"r"`
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
