package events

import "keesvv/go-tcp-chat/internals/message"

// Type represents an event type identifier.
type Type int

const (
	MessageType Type = 1
)

// Event represents a chat event.
type Event struct {
	Type Type `json:"t"`
}

// BaseEvent TODO
type BaseEvent struct {
	Event *Event `json:"e"`
}

// MessageEvent TODO
type MessageEvent struct {
	Event   *Event           `json:"e"`
	Message *message.Message `json:"msg"`
}

func newEvent(t Type) *Event {
	return &Event{
		Type: t,
	}
}

// NewMessageEvent TODO
func NewMessageEvent(msg *message.Message) *MessageEvent {
	return &MessageEvent{
		Event:   newEvent(MessageType),
		Message: msg,
	}
}
