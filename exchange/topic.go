package exchange

import (
	"github.com/verloop/daakia/shared"
	"strings"
	"sync"
)


type Topic interface {
	Subscribe(key[]byte, s *shared.Entity)
	//UnSubscribe(key[]byte, s *shared.Entity)
	Publish(key []byte, data []byte)
	//Remove(key[]byte)
}

type TopicConf struct {
	Separator string
	SingleWc string
	MultiWc string
	Sys string
	Dist func() shared.Distribution
}

type Node struct {
	Conf *TopicConf
	SubscriberKeys map[string]bool
	subscribers shared.Distribution
	children map[string]*Node
	mu sync.Mutex
}

func (n *Node) Publish(key []byte, data []byte) {
	// Bad implementation first :)
	arr := strings.Split(string(key),n.Conf.Separator)
	// It is us!!! :D
	if len(arr) == 1 {
		if n.subscribers == nil {
			// lolmax why write?
			return
		}
		n.subscribers.Write(data)
		return
	}

	if n.children == nil {
		n.children = make(map[string]*Node)
	}
	if v,ok:= n.children[arr[1]]; ok {
		v.Publish([]byte(strings.Join(arr[1:],n.Conf.Separator)),data)
	}

	return
}

func(n *Node) Subscribe(key[]byte, s *shared.Entity) {
	//Bad implementation
	arr := strings.Split(string(key),n.Conf.Separator)
	//It is us!
	if len(arr) == 1 || arr[1] == n.Conf.MultiWc {
		if n.subscribers == nil {
			n.subscribers = n.Conf.Dist()
		}

		if n.SubscriberKeys == nil {
			n.SubscriberKeys = make(map[string]bool)
		}
		if !n.SubscriberKeys[s.Id] {
			n.subscribers.Attach(s)
			n.SubscriberKeys[s.Id] = true
		}
		return
	}

	if arr[1] == n.Conf.SingleWc {
		k := []byte(strings.Join(arr[2:], n.Conf.Separator))
		n.mu.Lock()
		defer n.mu.Unlock()
		for _,v := range n.children {
			v.Subscribe(k, s)
		}
		return
	}
	n.mu.Lock()
	if n.children == nil {
		n.children = make(map[string]*Node)
	}
	if _,ok:=n.children[arr[1]]; !ok {
		n.children[arr[1]] = &Node{Conf:n.Conf}
	}
	n.mu.Unlock()
	n.children[arr[1]].Subscribe([]byte(strings.Join(arr[1:], n.Conf.Separator)), s)
}



