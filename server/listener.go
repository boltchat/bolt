package server

import (
	"net"

	"github.com/bolt-chat/server/handlers"
	"github.com/bolt-chat/server/logging"
)

// Listener TODO
type Listener struct {
	Bind []string
	Port int
}

func handleListener(conns []*net.TCPConn, l *net.TCPListener) error {
	for {
		conn, err := l.AcceptTCP()
		conns = append(conns, conn)
		// fmt.Println(conns)

		if err != nil {
			return err
		}

		// Accept new connection
		go handlers.HandleConnection(conns, conn)
	}
}

/*
Listen starts a new server/listener.
*/
func (listener *Listener) Listen() error {
	// All connections for this listener
	conns := make([]*net.TCPConn, 0, 5)

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

		go handleListener(conns, l)
	}

	return nil
}
