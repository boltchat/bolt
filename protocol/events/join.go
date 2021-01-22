package events

import (
	"keesvv/go-tcp-chat/protocol"
)

type JoinEvent struct {
	BaseEvent
	User *protocol.User `json:"user"`
}

func NewJoinEvent(user *protocol.User) *JoinEvent {
	return &JoinEvent{
		BaseEvent: BaseEvent{
			Event: newEvent(JoinType),
		},
		User: user,
	}
}
