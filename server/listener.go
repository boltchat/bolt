package server

import (
	"net"

	"github.com/bolt-chat/server/handlers"
	"github.com/bolt-chat/server/logging"
	"github.com/bolt-chat/util"
)

// Listener TODO
type Listener struct {
	Bind []string
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

		go handleListener(&connPool, l)
	}

	return nil
}
