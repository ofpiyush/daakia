package daakia

import (
	"github.com/daakia/daakia-go"
	"log"
)

const (
	Echo_Server_yolo = 11
)

type EchoClient interface {
}

func NewEchoClientRouter(client EchoClient, logger *log.Logger) *EchoRouter {
	return &EchoRouter {
		Client: client,
		Logger: logger,
	}
}

type EchoRouter struct {
	Client      EchoClient
	Logger      *log.Logger
}

func (r *EchoRouter) Route(buf []byte) error {
	switch buf[0] {
	}
	return nil
}

type EchoServer struct {
	Connection daakia.Conn
}

func(c *EchoServer) yolo() error {
	routing := []byte{Echo_Server_yolo}
	err := c.Connection.Send(routing)
	return err
}
