package format

import (
	"encoding/json"
	"fmt"
	"keesvv/go-tcp-chat/protocol/events"
	"strings"
	"time"

	"github.com/fatih/color"
)

func FormatMessage(e *events.BaseEvent) string {
	msgEvt := &events.MessageEvent{}
	json.Unmarshal(*e.Raw, msgEvt)

	sentAt := time.Unix(msgEvt.Message.SentAt, 0)

	timestamp := strings.Join([]string{
		color.HiBlackString("["),
		sentAt.Format(time.Stamp),
		color.HiBlackString("]"),
	}, "")

	return fmt.Sprintf(
		"%s <%s> %s",
		timestamp,
		msgEvt.Message.User.Nickname,
		msgEvt.Message.Content,
	)
}
