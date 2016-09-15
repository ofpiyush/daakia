package listeners

import (
	"log"
	"net"
	"time"
)

type TcpListener struct {
	handler func(net.Conn) error
	Logger  *log.Logger
}

func NewTCP() *TcpListener {
	return new(TcpListener)
}

func (s *TcpListener) Listen(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	var tempDelay time.Duration
	for {
		con, err := ln.Accept()
		// Lifted from net/http package
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				s.Logger.Printf("tcp: Accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return err
		}
		tempDelay = 0
		go s.handler(con)
	}

}

func (s *TcpListener) Handle(uri string, h func(net.Conn) error) {
	s.handler = h
}
