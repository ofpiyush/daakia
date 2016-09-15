package auth

import (
	"errors"
	"github.com/verloop/daakia/shared"
)

var ErrAlreadyAttached = errors.New("Entity Manager is already attached")


type Authorization interface {
	AttachEntityManager(shared.EntityManager) error
	Authorize([]byte) (*shared.Entity, error)
}
