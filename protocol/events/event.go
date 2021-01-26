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
