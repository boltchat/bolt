package format

import (
	"keesvv/go-tcp-chat/protocol/events"
)

type FormatHandler = func(e *events.BaseEvent) string
