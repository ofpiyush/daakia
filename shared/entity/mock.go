package entity

import (
	"github.com/verloop/daakia/shared"
)

type MockManager struct {
	Distribution func() shared.Distribution
	entities     map[string]*shared.Entity
}

func (m *MockManager) Get(id []byte) (*shared.Entity, error) {
	if m.entities == nil {
		m.entities = make(map[string]*shared.Entity)
	}
	if m.Distribution == nil {
		panic("Yo! Mock Entity manager needs a distribution")
	}
	if v, ok := m.entities[string(id)]; ok {
		return v, nil
	}
	m.entities[string(id)] = &shared.Entity{
		Id:   string(id),
		Dist: m.Distribution(),
	}
	m.entities[string(id)].Init()
	return m.entities[string(id)], nil
}

func (m *MockManager) Create(id []byte) (*shared.Entity, error) {
	return m.Get(id)
}
