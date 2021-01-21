package format

import (
	"keesvv/go-tcp-chat/internals/protocol/events"
)

type FormatHandler = func(e *events.BaseEvent) string
