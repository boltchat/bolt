package tui

import (
	"encoding/json"
	"fmt"
	"keesvv/go-tcp-chat/internals/protocol/events"
	"time"
)

type formatHandler = func(e *events.BaseEvent) string

func formatMessage(e *events.BaseEvent) string {
	msgEvt := &events.MessageEvent{}
	json.Unmarshal(*e.Raw, msgEvt)

	sentAt := time.Unix(msgEvt.Message.SentAt, 0)

	return fmt.Sprintf(
		"[%v] <%s> %s\n",
		sentAt.Format(time.Stamp),
		msgEvt.Message.User.Nickname,
		msgEvt.Message.Content,
	)
}
