package client

import (
	"net"

	"github.com/bolt-chat/protocol"
	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/util"
)

type Client struct {
	Conn *net.TCPConn // TODO: make private
	User protocol.User
	Opts Options
}

func NewClient(opts Options) *Client {
	return &Client{
		User: protocol.User{
			Nickname: opts.Nickname,
		},
		Opts: opts,
	}
}

// TODO: refactor
func (c *Client) Connect() error {
	var ip net.IP
	var port int = c.Opts.Port

	if parsedIP := net.ParseIP(c.Opts.Hostname); parsedIP != nil {
		ip = parsedIP
	}

	if ip == nil {
		_, srvAddrs, _ := net.LookupSRV("bolt", "tcp", c.Opts.Hostname)

		if len(srvAddrs) > 0 {
			targetIps, _ := net.LookupIP(srvAddrs[0].Target)
			ip = targetIps[0]
			port = int(srvAddrs[0].Port)
		}
	}

	if ip == nil {
		ips, lookupErr := net.LookupIP(c.Opts.Hostname)
		if lookupErr != nil {
			return lookupErr
		}

		ip = ips[0]
	}

	conn, dialErr := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   ip,
		Port: port,
	})

	if dialErr != nil {
		return dialErr
	}

	// Set the connection
	c.Conn = conn

	util.WriteJson(conn, *events.NewJoinEvent(&c.User))
	return nil
}
