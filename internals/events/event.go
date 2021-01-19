package events

// Type represents an event type identifier.
type Type int

const (
	// MessageEvent represents a message event.
	MessageEvent Type = iota
)

// Event represents a chat event.
type Event struct {
	Type Type `json:"t"`
}
