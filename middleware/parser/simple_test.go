package parser

import (
	"encoding/binary"
	"io"
	"testing"
	"math/rand"
	"bytes"
)
var (
	_ = bytes.ErrTooLarge
)

type TestRouter struct {
	Expected []byte
	T *testing.T
}

func (t *TestRouter) Route(in []byte) {
	if t.T !=nil && !bytes.Equal(in, t.Expected) {
		t.T.Errorf("Mismatched bytes, expected %d, got %d", len(t.Expected), len(in))
	}
}

type MsgGenerator struct {
	message []byte
	Length int
	Times int
}

func (m *MsgGenerator) Generate() []byte {
	m.Length +=4
	m.message = make([]byte,m.Length)
	binary.LittleEndian.PutUint32(m.message[:4],uint32(m.Length-4))
	rand.Read(m.message[4:])
	return m.message
	//m.message = append(m.message,m.message...)
}

func (m *MsgGenerator) Read(p []byte) (int, error) {
	if m.Times == 0 {
		return 0,io.EOF
	}
	m.Times--
	return copy(p,m.message), nil
}

func (m *MsgGenerator) Close() error{ return nil;}


func TestSimple_Parse(t *testing.T) {
	base := "hey "
	str_msg := ""
	for i := 0; i < 1000; i++ {
		str_msg += base
	}


	buf := make([]byte, len(str_msg)+4, len(str_msg)+4)

	binary.LittleEndian.PutUint32(buf[:4], uint32(len(str_msg)))
	copy(buf[4:len(str_msg)+4], []byte(str_msg)[:len(str_msg)])

	tr := &TestRouter{
		Expected: buf,
		T: t,
	}
	s := &Simple{}
	r, w := io.Pipe()

	two_bufs := make([]byte, len(buf)*2, len(buf)*2)
	copy(two_bufs, buf)
	copy(two_bufs[len(buf):], buf)
	go s.Parse(r,tr)
	//w.Write(buf)
	//w.Write(two_bufs)
	two_n_5 := make([]byte, len(two_bufs)+5, len(two_bufs)+5)
	copy(two_n_5, two_bufs)
	copy(two_n_5[len(two_bufs):], buf[:5])
	//w.Write(two_n_5)
	// Write rest half
	//w.Write(buf[5:])

	two_n_1 := make([]byte, len(two_bufs)+1, len(two_bufs)+1)
	copy(two_n_1, two_bufs)
	copy(two_n_1[len(two_bufs):], buf[:1])
	w.Write(two_n_1)
	w.Write(buf[1:])
	t.Log(s.processed)
}

func Benchmark__Simple_Parse(b *testing.B) {
	b.StopTimer()

	s := &Simple{}
	m := &MsgGenerator{Length:100,Times:b.N}
	tr := &TestRouter{
		Expected: m.Generate(),
	}
	b.SetBytes(int64(m.Length))

	b.StartTimer()

	s.Parse(m,tr)

	b.StopTimer()
}
