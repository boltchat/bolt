package events

import "keesvv/go-tcp-chat/internals/message"

// MessageEvent TODO
type MessageEvent struct {
	BaseEvent
	Message *message.Message `json:"msg"`
}

// NewMessageEvent TODO
func NewMessageEvent(msg *message.Message) *MessageEvent {
	return &MessageEvent{
		BaseEvent: BaseEvent{
			Event: newEvent(MessageType),
		},
		Message: msg,
	}
}
