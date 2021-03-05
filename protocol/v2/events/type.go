package events

type EventType uint8

const (
	// JoinEvent is the event type used for new connections.
	JoinEvent EventType = 0x01
	// LeaveEvent is the event type used for disconnects.
	LeaveEvent EventType = 0x02
	// MessageEvent is the event type used for messages.
	MessageEvent EventType = 0x03
	// CommandEvent is the event type used for commands.
	CommandEvent EventType = 0x04
	// ErrorEvent is the event type used for error reporting.
	ErrorEvent EventType = 0x05
	// MOTDEvent is the event type used for the Message-of-the-Day (MOTD).
	MOTDEvent EventType = 0x06
	// NoticeEvent is the event type used for general server notices.
	NoticeEvent EventType = 0x07
	// TypingEvent is the event type used for typing indicators.
	TypingEvent EventType = 0x08
)
