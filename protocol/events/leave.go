package events

import (
	"github.com/bolt-chat/protocol"
)

type LeaveEvent struct {
	BaseEvent
	User *protocol.User `json:"user"`
}

func NewLeaveEvent(user *protocol.User) *LeaveEvent {
	return &LeaveEvent{
		BaseEvent: *NewBaseEvent(LeaveType),
		User:      user,
	}
}
