package client

import (
	"bytes"
	"encoding/json"
	"keesvv/go-tcp-chat/protocol/events"
)

func (c *Connection) ReadEvents(evts chan *events.BaseEvent) error {
	for {
		b := make([]byte, 4096)
		_, err := c.TCPConn.Read(b)

		if err != nil {
			return err
		}

		b = bytes.TrimRight(b, "\x00")
		evt := &events.BaseEvent{}
		jsonErr := json.Unmarshal(b, evt)

		evt.Raw = &b

		if jsonErr != nil {
			return err
		}

		go func() {
			evts <- evt
		}()
	}
}
