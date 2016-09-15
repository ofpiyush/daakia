package distribution

import (
	"container/ring"
	"github.com/verloop/daakia/shared"
	"io"
	"sync"
	"fmt"
)


func NewRoundRobin() shared.Distribution {
	return &RoundRobin{
		connections: ring.New(1),
	}
}

type RoundRobin struct {
	connections *ring.Ring
	mu          sync.Mutex
}

func (r *RoundRobin) Attach(c io.WriteCloser) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	nr := ring.New(1)
	nr.Value = c
	if r.connections == nil {
		r.connections = nr
		return nil
	}
	r.connections.Link(nr)
	return nil
}

func (r *RoundRobin) Write(b []byte) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.connections != nil {
		v := r.connections
		for v.Value == nil {
			v = r.connections.Unlink(1)
		}
		for _, err := v.Value.(io.WriteCloser).Write(b); err != nil; {
			fmt.Println(err)
			v.Value = nil
			v = v.Next()
		}
		r.connections = v
	}
	return len(b), nil
}
