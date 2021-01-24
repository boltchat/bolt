package events

type MotdEvent struct {
	BaseEvent
	Motd string
}

func NewMotdEvent(motd string) *MotdEvent {
	return &MotdEvent{
		BaseEvent: *NewBaseEvent(MotdType),
		Motd:      motd,
	}
}
