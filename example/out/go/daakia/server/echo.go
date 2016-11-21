package daakia

import (
	"github.com/daakia/daakia-go"
	"log"
)

const (
	Echo_Server_yolo = 11
)

type EchoServer interface {
	yolo() error
}

func NewEchoServerRouter(server EchoServer, logger *log.Logger) *EchoRouter {
	return &EchoRouter {
		Server: server,
		Logger: logger,
	}
}

type EchoRouter struct {
	Server      EchoServer
	Logger      *log.Logger
}

func (r *EchoRouter) Route(buf []byte) error {
	switch buf[0] {
	case Echo_Server_yolo:
		err := r.Server.yolo()
		if err != nil {
			r.Logger.Println(err)
			return err
		}
	}
	return nil
}

type EchoClient struct {
	Connection daakia.Conn
}
