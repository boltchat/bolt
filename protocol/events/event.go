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

// Event represents a chat event.
type Event struct {
	Type      Type  `json:"t"`
	CreatedAt int64 `json:"c"`
}

// BaseEvent TODO
type BaseEvent struct {
	Event *Event `json:"e"`
	Raw   *[]byte
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
