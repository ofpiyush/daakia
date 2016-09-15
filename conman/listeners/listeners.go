package listeners

import (
	"net"
)

// Try to be as close to http and websockets
type Listener interface {
	Handle(string, func(net.Conn) error)
	Listen(string) error
	// Shutdown()
}
