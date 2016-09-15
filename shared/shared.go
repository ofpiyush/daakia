package shared

import (
	"io"
	"net"
	"sync"
)

type Distribution interface {
	Attach(io.WriteCloser) error
	Write([]byte) (int, error)
}

type EntityManager interface {
	Get([]byte) (*Entity, error)
	Create([]byte) (*Entity, error)
}

type Entity struct {
	Id   string
	Dist Distribution
	r    io.Reader
	w    io.Writer
	mu   sync.Mutex
}

func (e *Entity) Init() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.r, e.w = io.Pipe()
}

// Attaches a net.Conn to the underlying Distribution channel and sets up reading
// PS: Will break if Distribution is not setup.
// Unless you really know what you are doing, always use an entity manager
func (e *Entity) AttachConn(c net.Conn) error {
	e.mu.Lock()
	e.mu.Unlock()
	// Keep writing to internal pipe writer
	go func(r io.ReadCloser) {
		io.Copy(e.w, r)
	}(c)
	return e.Dist.Attach(c)
}

func (e *Entity) Read(p []byte) (int, error) {
	return e.r.Read(p)
}

// Implements the io.Writer interface, writes to the underlying Distribution
// PS: Will break if Distribution is not setup.
// Unless you really know what you are doing, always use an entity manager
func (e *Entity) Write(p []byte) (int, error) {

	return e.Dist.Write(p)
}


func(e *Entity) Close() error {
	// Maybe support closing all connections of a user?
	// Laters baby
	return nil
}