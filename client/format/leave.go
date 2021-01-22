package format

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	"github.com/gdamore/tcell/v2"
	"github.com/keesvv/bolt.chat/protocol/events"
)

func FormatLeave(e *events.BaseEvent) string {
	leaveEvt := &events.JoinEvent{}
	json.Unmarshal(*e.Raw, leaveEvt)

	return color.HiMagentaString(
		fmt.Sprintf("%s %s left the room.", string(tcell.RuneDiamond), leaveEvt.User.Nickname),
	)
}
