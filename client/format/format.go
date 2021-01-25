package format

import (
	"github.com/bolt-chat/protocol/events"
)

type FormatHandler = func(e *events.BaseEvent) string

var FormatMap = map[events.Type]FormatHandler{
	events.MotdType:    FormatMotd,
	events.MessageType: FormatMessage,
	events.ErrorType:   FormatError,
	events.JoinType:    FormatJoin,
	events.LeaveType:   FormatLeave,
}
