package events

// Type represents an event type identifier.
type Type int

const (
	MessageType Type = iota
)

// Event represents a chat event.
type Event struct {
	Type Type `json:"t"`
}

// BaseEvent TODO
type BaseEvent struct {
	Event *Event `json:"e"`
}

func newEvent(t Type) *Event {
	return &Event{
		Type: t,
	}
}
