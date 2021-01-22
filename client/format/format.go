package format

import (
	"keesvv/bolt.chat/protocol/events"
)

type FormatHandler = func(e *events.BaseEvent) string
