package daakia

import (
	"github.com/daakia/daakia-go"
	"log"
)

const (
	TwoWayPublish_Server_Publish = 11
	TwoWayPublish_Server_PubAck = 12
	TwoWayPublish_Client_Publish = 13
	TwoWayPublish_Client_PubAck = 14
)

type TwoWayPublishServer interface {
	Publish([]byte) error
	PubAck([]byte) error
}

func NewTwoWayPublishServerRouter(server TwoWayPublishServer, logger *log.Logger) *TwoWayPublishRouter {
	return &TwoWayPublishRouter {
		Server: server,
		Logger: logger,
	}
}

type TwoWayPublishRouter struct {
	Server      TwoWayPublishServer
	Logger      *log.Logger
}

func (r *TwoWayPublishRouter) Route(buf []byte) error {
	switch buf[0] {
	case TwoWayPublish_Server_Publish:
		err := r.Server.Publish(buf[1:])
		if err != nil {
			r.Logger.Println(err)
			return err
		}
	case TwoWayPublish_Server_PubAck:
		err := r.Server.PubAck(buf[1:])
		if err != nil {
			r.Logger.Println(err)
			return err
		}
	}
	return nil
}

type TwoWayPublishClient struct {
	Connection daakia.Conn
}

func(c *TwoWayPublishClient) Publish(payload []byte) error {
	routing := []byte{TwoWayPublish_Client_Publish}
	err := c.Connection.Send(routing, payload)
	return err
}

func(c *TwoWayPublishClient) PubAck(payload []byte) error {
	routing := []byte{TwoWayPublish_Client_PubAck}
	err := c.Connection.Send(routing, payload)
	return err
}
