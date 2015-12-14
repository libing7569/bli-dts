package communicate

import "net"

type CommunicateEntity struct {
	net         string
	addr        string
	rawHandler  func(net.Conn, chan<- []byte) error
	dataHandler func(<-chan []byte) error
}
