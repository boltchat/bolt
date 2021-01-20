package events

import "keesvv/go-tcp-chat/internals/protocol"

// MessageEvent TODO
type MessageEvent struct {
	BaseEvent
	Message *protocol.Message `json:"msg"`
}

// NewMessageEvent TODO
func NewMessageEvent(msg *protocol.Message) *MessageEvent {
	return &MessageEvent{
		BaseEvent: BaseEvent{
			Event: newEvent(MessageType),
		},
		Message: msg,
	}
}
