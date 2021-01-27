// bolt.chat
// Copyright (C) 2021  The bolt.chat Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
		logging.LogListener(l.Addr().String())

		go handleListener(&connPool, l)
	}

	return nil
}
