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

type TwoWayPublishClient interface {
	Publish([]byte) error
	PubAck([]byte) error
}

func NewTwoWayPublishClientRouter(client TwoWayPublishClient, logger *log.Logger) *TwoWayPublishRouter {
	return &TwoWayPublishRouter {
		Client: client,
		Logger: logger,
	}
}

type TwoWayPublishRouter struct {
	Client      TwoWayPublishClient
	Logger      *log.Logger
}

func (r *TwoWayPublishRouter) Route(buf []byte) error {
	switch buf[0] {
	case TwoWayPublish_Client_Publish:
		err := r.Client.Publish(buf[1:])
		if err != nil {
			r.Logger.Println(err)
			return err
		}
	case TwoWayPublish_Client_PubAck:
		err := r.Client.PubAck(buf[1:])
		if err != nil {
			r.Logger.Println(err)
			return err
		}
	}
	return nil
}

type TwoWayPublishServer struct {
	Connection daakia.Conn
}

func(c *TwoWayPublishServer) Publish(payload []byte) error {
	routing := []byte{TwoWayPublish_Server_Publish}
	err := c.Connection.Send(routing, payload)
	return err
}

func(c *TwoWayPublishServer) PubAck(payload []byte) error {
	routing := []byte{TwoWayPublish_Server_PubAck}
	err := c.Connection.Send(routing, payload)
	return err
}
