package server

import (
	"keesvv/go-tcp-chat/internals/handlers"
	"net"
)

type Listener struct {
	IP   string
	Port int
}

/*
Listen starts a new server/listener.
*/
func (listener *Listener) Listen() error {
	l, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.ParseIP(listener.IP),
		Port: listener.Port,
	})

	if err != nil {
		return err
	}

	for {
		conn, err := l.AcceptTCP()

		if err != nil {
			return err
		}

		// Accept new connection
		go handlers.HandleConnection(conn)
	}
}
