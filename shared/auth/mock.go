package auth

import "github.com/verloop/daakia/shared"



type MockAuth struct {
	em shared.EntityManager
}

func (m *MockAuth) AttachEntityManager(em shared.EntityManager) error {
	if m.em != nil {
		return ErrAlreadyAttached
	}
	m.em = em
	return nil
}

func (m *MockAuth) Authorize(u []byte) (*shared.Entity, error) {
	return m.em.Get(u)
}
