package events

import "keesvv/go-tcp-chat/internals/user"

type JoinEvent struct {
	BaseEvent
	User *user.User `json:"user"`
}

func NewJoinEvent(user *user.User) *JoinEvent {
	return &JoinEvent{
		BaseEvent: BaseEvent{
			Event: newEvent(JoinType),
		},
		User: user,
	}
}
