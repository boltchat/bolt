package events

type Event struct {
	Header     *Header
	CRC32      uint32
	Signature  *[]byte
	RawPayload *[]byte
}

// NewEvent TODO
func NewEvent(t EventType, payload *[]byte) *Event {
	return &Event{
		Header: &Header{ // TODO:
			Version:   1,
			EventType: t,
		},
		RawPayload: payload,
	}
}
