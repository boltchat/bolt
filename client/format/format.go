package format

import (
	"github.com/keesvv/bolt.chat/protocol/events"
)

type FormatHandler = func(e *events.BaseEvent) string
