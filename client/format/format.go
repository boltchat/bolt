package format

import (
	"github.com/bolt-chat/protocol/events"
)

type FormatHandler = func(e *events.BaseEvent) string
