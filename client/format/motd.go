package format

import (
	"encoding/json"
	"keesvv/bolt.chat/protocol/events"

	"github.com/fatih/color"
)

func FormatMotd(e *events.BaseEvent) string {
	motdEvt := &events.MotdEvent{}
	json.Unmarshal(*e.Raw, motdEvt)

	return color.HiCyanString(motdEvt.Motd)
}
