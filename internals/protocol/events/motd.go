package events

type MotdEvent struct {
	BaseEvent
	Motd string
}

func NewMotdEvent(motd string) *MotdEvent {
	return &MotdEvent{
		BaseEvent: BaseEvent{
			Event: newEvent(MotdType),
		},
		Motd: motd,
	}
}
