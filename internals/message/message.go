package message

import "net"

type Message struct {
	Content string
}

func (m *Message) Send(conn *net.TCPConn) {
	conn.Write([]byte(m.Content)) // TODO:
}
