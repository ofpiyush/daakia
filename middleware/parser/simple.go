package parser

import (
	"io"
	"encoding/binary"
	"errors"
	"github.com/daakia/middleware"
)

var (
	ErrNoNext = errors.New("No next to call")
)

//Todo: Move this inside TCP as websocket already handles these internally.
type Simple struct {
	processed       uint32
	prev            uint32
	carry, i, lc, e uint32
	ln              uint32
	nxt             []byte
	router 			middleware.Router
}
// Use channels for now, improve later
func (s *Simple) Parse(r io.ReadCloser, router middleware.Router) error {
	buf := make([]byte, 4096)
	s.router = router
	if s.router == nil {
		return ErrNoNext
	}
	for {
		n, err := r.Read(buf[s.carry:])
		if err != nil {
			r.Close()
			return err
		}
		if n == 0 {
			continue
		}

		err = s.parse(buf, s.carry + uint32(n))
		if err != nil {
			return err
		}
	}
}

func (s *Simple) parse(buf []byte, n uint32) (err error) {
	s.carry = 0
	if s.prev > 0 {
		if n < s.prev {
			copy(s.nxt[s.lc:], buf[:n])
			s.prev -= n
			s.lc = s.lc + n
			return
		}
		copy(s.nxt[s.lc:], buf[:s.prev])

		s.router.Route(s.nxt)
	}
	for s.i = s.prev; s.i < n; {
		if n-s.i < 4 {
			// Copy current int into carry
			s.carry = n - s.i
			// Move current bits to the start
			copy(buf[0:], buf[s.i:n])
			s.prev = 0
			return
		}

		s.ln = binary.LittleEndian.Uint32(buf[s.i:s.i+4])

		s.e = s.i + s.ln +4
		if s.e > n {
			s.e = n
		}
		s.nxt = make([]byte, s.ln+4, s.ln+4)
		s.lc = uint32(copy(s.nxt, buf[s.i:s.e]))
		s.i += s.ln +4
		// Use a ring or even channels later
		if s.i <= n {
			s.router.Route(s.nxt)
		}
	}
	if s.i > n {
		s.prev = s.i - n
		return
	}
	s.prev = 0
	return
}