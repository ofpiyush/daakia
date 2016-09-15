package conman

import (
	"github.com/verloop/daakia/conman/listeners"
)

type Server struct {
	protocol  Protocol
	listeners map[string]listeners.Listener
}

func (s *Server) AttachListener(addr string, listener listeners.Listener) {
	if s.listeners == nil {
		s.listeners = make(map[string]listeners.Listener)
	}
	s.listeners[addr] = listener
}

func (s *Server) ListenAndServe() error {
	err := make(chan error)

	for port, listener := range s.listeners {
		go func(p string, l listeners.Listener) {
			err <- l.Listen(p)
		}(port, listener)
	}
	for {
		select {
		case x := <-err:
			if x != nil {
				return x
			}
		}
	}
	return nil
}
