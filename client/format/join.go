package format

import (
	"encoding/json"
	"fmt"

	"github.com/bolt-chat/protocol/events"
	"github.com/fatih/color"
	"github.com/gdamore/tcell/v2"
)

func FormatJoin(e *events.BaseEvent) string {
	joinEvt := &events.JoinEvent{}
	json.Unmarshal(*e.Raw, joinEvt)

	return color.HiMagentaString(
		fmt.Sprintf("%s %s joined the room.", string(tcell.RuneDiamond), joinEvt.User.Nickname),
	)
}
