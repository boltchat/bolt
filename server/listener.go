package server

import (
	"net"

	"github.com/bolt-chat/server/handlers"
	"github.com/bolt-chat/server/logging"
	"github.com/bolt-chat/util"
)

type Bind struct {
	Address string
	Proto   string
}

// Listener TODO
type Listener struct {
	Bind []Bind
	Port int
}

/*
handleListener handles an individual TCP listener.
*/
func handleListener(pool *util.ConnPool, l *net.TCPListener) error {
	for {
		conn, err := l.AcceptTCP()
		pool.AddToPool(conn)

		if err != nil {
			return err
		}

		// Accept new connection
		go handlers.HandleConnection(pool, conn)
	}
}

/*
Listen starts a new server/listener.
*/
func (listener *Listener) Listen() error {
	// The connection pool for this listener
	connPool := make(util.ConnPool, 0, 5)

	for _, bind := range listener.Bind {
		l, err := net.ListenTCP(bind.Proto, &net.TCPAddr{
			IP:   net.ParseIP(bind.Address),
			Port: listener.Port,
		})

		if err != nil {
			return err
		}

		// TODO
		logging.LogListener(bind.Address, listener.Port)

		go handleListener(&connPool, l)
	}

	return nil
}
