package client

import (
	"bytes"
	"encoding/json"
	"keesvv/go-tcp-chat/protocol"
	"keesvv/go-tcp-chat/protocol/events"
	"keesvv/go-tcp-chat/util"
	"net"
)

func Connect(opts Options) (*Connection, error) {
	ips, lookupErr := net.LookupIP(opts.Hostname)
	if lookupErr != nil {
		return &Connection{}, lookupErr
	}

	ip := ips[0]

	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   ip,
		Port: opts.Port,
	})

	if err != nil {
		return &Connection{}, err
	}

	user := &protocol.User{
		Nickname: opts.Nickname,
	}

	util.WriteJson(conn, *events.NewJoinEvent(user))

	return &Connection{
		TCPConn: conn,
		User:    *user,
	}, nil
}

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

		// switch evt.Event.Type {
		// case events.MessageType:
		// 	msgEvt := &events.MessageEvent{}
		// 	json.Unmarshal(b, msgEvt)
		// 	go func() { evts <- msgEvt.Message.Content }() // TODO
		// case events.MotdType:
		// 	motdEvt := &events.MotdEvent{}
		// 	json.Unmarshal(b, motdEvt)
		// 	go func() { evts <- motdEvt.Motd }()
		// case events.ErrorType:
		// 	errEvt := &events.ErrorEvent{}
		// 	json.Unmarshal(b, errEvt)
		// 	fmt.Println(errEvt.Error) // TODO
		// }
	}
}
