package events

type Event struct {
	Header     *Header
	CRC32      uint32
	Signature  *[]byte
	RawPayload *[]byte
	Payload    interface{}
}

// NewEvent TODO
func NewEvent(t EventType, payload interface{}) *Event {
	return &Event{
		Header: &Header{ // TODO:
			Version:   1,
			EventType: t,
		},
		Payload: payload,
	}
}
