package server

import (
	"crypto/tls"
	"fmt"
	"net"

	"github.com/keesvv/bolt.chat/server/handlers"
	"github.com/keesvv/bolt.chat/server/logging"
)

// Listener TODO
type Listener struct {
	Bind []string
	Port int
}

func handleListener(conns []net.Conn, l net.Listener) error {
	for {
		conn, err := l.Accept()
		conns = append(conns, conn)

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
	conns := make([]net.Conn, 0, 5)

	cert, certErr := tls.LoadX509KeyPair("server.crt", "server.key") // TODO:

	if certErr != nil {
		panic(certErr)
	}

	for _, ip := range listener.Bind {
		l, err := tls.Listen("tcp", fmt.Sprintf("%s:%s", ip, listener.Port), &tls.Config{
			Certificates: []tls.Certificate{
				cert,
			},
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
