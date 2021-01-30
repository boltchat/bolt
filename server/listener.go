// Copyright 2021 The boltchat Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"net"

	"github.com/bolt-chat/server/handlers"
	"github.com/bolt-chat/server/logging"
	"github.com/bolt-chat/server/pools"
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
func handleListener(pool *pools.ConnPool, l *net.TCPListener) error {
	for {
		tcpConn, err := l.AcceptTCP()
		conn := pools.NewConnection(tcpConn, nil)

		// Add connection to pool
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
	connPool := make(pools.ConnPool, 0, 5)

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
