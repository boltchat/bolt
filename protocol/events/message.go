package events

import "github.com/bolt-chat/protocol"

// MessageEvent TODO
type MessageEvent struct {
	BaseEvent
	Message *protocol.Message `json:"msg"`
}

// NewMessageEvent TODO
func NewMessageEvent(msg *protocol.Message) *MessageEvent {
	return &MessageEvent{
		BaseEvent: *NewBaseEvent(MessageType),
		Message:   msg,
	}
}
