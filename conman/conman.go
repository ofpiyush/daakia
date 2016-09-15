package conman

import (
	"net"
)

type Protocol interface {
	NewConnection(con net.Conn) error
	GetPath() string
}
