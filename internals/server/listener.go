package server

import (
	"keesvv/go-tcp-chat/internals/handlers"
	"keesvv/go-tcp-chat/internals/logging"
	"net"
)

// Listener TODO
type Listener struct {
	Bind []string
	Port int
}

func handleListener(l *net.TCPListener) error {
	for {
		conn, err := l.AcceptTCP()

		if err != nil {
			return err
		}

		// Accept new connection
		go handlers.HandleConnection(conn)
	}
}

/*
Listen starts a new server/listener.
*/
func (listener *Listener) Listen() error {
	for _, ip := range listener.Bind {
		l, err := net.ListenTCP("tcp", &net.TCPAddr{
			IP:   net.ParseIP(ip),
			Port: listener.Port,
		})

		if err != nil {
			return err
		}

		// TODO
		logging.LogListener(ip, listener.Port)

		go handleListener(l)
	}

	return nil
}
