package events

type EventType int

const (
	JoinEvent    EventType = 0x01
	LeaveEvent   EventType = 0x02
	MessageEvent EventType = 0x03
	CommandEvent EventType = 0x04
	ErrorEvent   EventType = 0x05
	MOTDEvent    EventType = 0x06
	NoticeEvent  EventType = 0x07
	TypingEvent  EventType = 0x08
)
