package events

import (
	"github.com/bolt-chat/protocol"
)

type JoinEvent struct {
	BaseEvent
	User *protocol.User `json:"user"`
}

func NewJoinEvent(user *protocol.User) *JoinEvent {
	return &JoinEvent{
		BaseEvent: *NewBaseEvent(JoinType),
		User:      user,
	}
}
