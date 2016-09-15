package distribution

import (

	"container/ring"
	"github.com/verloop/daakia/shared"
	"io"

	"sync"
)

func NewFanOut() shared.Distribution {
	return &FanOut{
		connections: ring.New(1),
	}
}

type FanOut struct {
	connections *ring.Ring
	amu         sync.Mutex
	wmu         sync.Mutex
}

func (f *FanOut) Attach(c io.WriteCloser) error {
	f.amu.Lock()
	defer f.amu.Unlock()
	nr := ring.New(1)
	nr.Value = c
	f.connections.Link(nr)
	return nil
}

func (f *FanOut) Write(b []byte) (int, error) {
	f.wmu.Lock()
	defer f.wmu.Unlock()
	go func() {
		f.connections.Do(func(v interface{}) {
			if v != nil {
				v.(io.Writer).Write(b)
			}
		})
	}()
	return len(b), nil
}
