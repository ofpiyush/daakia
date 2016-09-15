package main

import (
	"fmt"
	"github.com/verloop/daakia/conman"
	"github.com/verloop/daakia/conman/listeners"
	"github.com/verloop/daakia/daak"
	"github.com/verloop/daakia/shared/distribution"
	"github.com/verloop/daakia/shared/entity"
	"net"
	"errors"
	"github.com/verloop/daakia/shared"
	"github.com/verloop/daakia/exchange"
)

type Daak struct {
	id *shared.Entity
	Manager shared.EntityManager
	Conn net.Conn
	Topic exchange.Topic
}

func(d *Daak) Unknown() error {
	d.Conn.Close()
	return errors.New("Phut le")
}

func(d *Daak) Ping() error {
	_,err := d.Conn.Write([]byte("Yo\n"))
	return err
}

func (d *Daak) Connect(id []byte) error {
	d.id ,_ = d.Manager.Get(id)
	d.id.AttachConn(d.Conn)
	return nil
}

func(d *Daak) Pub(to []byte, data []byte) error {
	if d.id == nil {
		d.Unknown()
	}
	to = []byte(fmt.Sprintf("/%s",to))
	d.Topic.Subscribe(to,d.id)
	d.Topic.Publish(to,[]byte(fmt.Sprintf("SAID %s %s\n", to, data)))
	return nil
}

func main() {
	distribution.Register("fanout", distribution.NewFanOut)
	distribution.Register("roundrobin", distribution.NewRoundRobin)

	d , err := distribution.Get("fanout")
	if err != nil {
		fmt.Println(err)
	}
	server := &conman.Server{}
	k := &entity.MockManager{Distribution: d}
	topic := &exchange.Node{
		Conf:&exchange.TopicConf{
			MultiWc: "#",
			SingleWc: "*",
			Separator: "/",
			Dist:d,
		},
	}

	ParseQ := func(n net.Conn) error {
		return daak.Parse(n, &Daak{
			Conn:n,
			Manager: k,
			Topic: topic,
		})
	}


	l := listeners.NewTCP()
	l.Handle("", ParseQ)
	server.AttachListener(":3000",l)

	l1 := listeners.NewWebSocket()
	l1.Handle("/mqtt", ParseQ)
	server.AttachListener(":8085", l1)

	// Fix http later
	//h := listeners.NewHTTP()
	//h.Handle("/mqtt->http://127.0.0.1:9000/lala", func() p.NewConnection)
	//server.AttachListener(":9091", h)

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}


