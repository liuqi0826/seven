package network

import (
	"net"
	"time"
)

const (
	NETWORK_TYPE_LOCAL      = "local"
	NETWORK_TYPE_TCP        = "tcp"
	NETWORK_TYPE_UDP        = "udp"
	NETWORK_TYPE_WEB_SOCKET = "webSocket"
)

type WebSocketConfig struct {
	URL    string
	SSLCRT string
	SSLKey string
}

func TCPConnect(n *Network, addr string) {
	for {
		conn, err := net.Dial("tcp", addr)
		if err == nil {
			connect := new(TCPConnection)
			connect.TCPConnection(conn)
			n.Create(connect)
			break
		} else {
			time.Sleep(time.Microsecond * 10)
		}
	}
}

//func LocalConnect(n *Network, l *Listener) {
//	c := new(ChanConnection)
//	c.ChanConnection(nil)
//	n.Create(c)
//	l.localChan <- c
//}
