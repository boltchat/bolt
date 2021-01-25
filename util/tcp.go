package util

import (
	"encoding/json"
	"net"

	"github.com/bolt-chat/server/logging"
)

type ConnPool []*net.TCPConn

func (c *ConnPool) AddToPool(conn *net.TCPConn) {
	// Append connection to pool
	*c = append(*c, conn)

	logging.LogDebug(
		"connection added to pool:",
		conn.RemoteAddr().String(),
	)

	logging.LogDebug("pool size:", len(*c))
}

func (c *ConnPool) RemoveFromPool(conn *net.TCPConn) {
	// Range through pool
	for i, curConn := range *c {
		// Target connection is found
		if curConn == conn {
			/*
				This removes the connection from the pool by its
				respective index.

				`i` represents the index of the matched connection.
				The first arg represents a slice of the pool that
				ends at the index of the connection in question.
				The second arg represents a slice of the pool that
				starts from the index of the connection in question + 1.
			*/
			*c = append((*c)[:i], (*c)[i+1:]...)
			return
		}
	}
}

func WriteJson(conn *net.TCPConn, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	conn.Write(b)
}

func Broadcast(conns *ConnPool, data interface{}) {
	for _, conn := range *conns {
		WriteJson(conn, data)
	}
}
