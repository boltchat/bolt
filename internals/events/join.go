package events

import "keesvv/go-tcp-chat/internals/user"

type JoinEvent struct {
	Event *Event     `json:"e"`
	User  *user.User `json:"user"`
}

func NewJoinEvent(user *user.User) *JoinEvent {
	return &JoinEvent{
		Event: newEvent(JoinType),
		User:  user,
	}
}
