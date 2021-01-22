package format

import (
	"encoding/json"
	"fmt"

	"github.com/keesvv/bolt.chat/protocol/events"

	"github.com/fatih/color"
)

func FormatError(e *events.BaseEvent) string {
	errEvt := &events.ErrorEvent{}
	json.Unmarshal(*e.Raw, errEvt)

	return color.HiRedString(
		fmt.Sprintf("[!] %s", errEvt.Error),
	)
}
