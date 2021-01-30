package client

import (
	"bytes"
	"encoding/json"

	"github.com/bolt-chat/client/errs"
	"github.com/bolt-chat/protocol/events"
)

func (c *Client) ReadEvents(evts chan *events.BaseEvent, closed chan bool) {
	for {
		// Allocate 64KB for the event
		// TODO: automatically resize
		b := make([]byte, 65536)
		_, err := c.Conn.Read(b)

		if err != nil {
			closed <- true
			return
		}

		b = bytes.TrimRight(b, "\x00")

		evt := &events.BaseEvent{}
		jsonErr := json.Unmarshal(b, evt)

		evt.Raw = &b

		if jsonErr != nil {
			errs.Emerg(err)
		}

		evts <- evt
	}
}
