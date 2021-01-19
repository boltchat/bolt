package events

import "keesvv/go-tcp-chat/internals/message"

// MessageEvent TODO
type MessageEvent struct {
	Event   *Event           `json:"e"`
	Message *message.Message `json:"msg"`
}

// NewMessageEvent TODO
func NewMessageEvent(msg *message.Message) *MessageEvent {
	return &MessageEvent{
		Event:   newEvent(MessageType),
		Message: msg,
	}
}
