package events

// ErrorEvent TODO
type ErrorEvent struct {
	BaseEvent
	Error string `json:"err"`
}

// NewErrorEvent TODO
func NewErrorEvent(err string) *ErrorEvent {
	return &ErrorEvent{
		BaseEvent: *NewBaseEvent(ErrorType),
		Error:     err,
	}
}
