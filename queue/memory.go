package queue

import (
	"sync"
	"time"
)

type Element struct {
	Data []byte
	TTL time.Time
	Next *Element
}

type Memory struct {
	mu sync.RWMutex
	len int64
	head *Element
	tail *Element
}

func(m *Memory) Push(d []byte, ttl time.Time) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tail.Next = &Element{Data:d, TTL:ttl}
	m.tail = m.tail.Next
}

func (m *Memory) Pop() ([]byte) {
	m.mu.RLock()
	defer func() {
		m.head, m.head.Next = m.head.Next,nil
		m.mu.RUnlock()
	}()
	for m.head {
		if m.head.TTL.After(time.Now()) {
			m.head = m.head.Next
		} else {
			break
		}
	}

	return m.head
}

func (m *Memory) Len() int64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.len
}

