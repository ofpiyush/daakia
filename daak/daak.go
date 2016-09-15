package daak

type Routes interface {
	Unknown() error
	Ping() error
	Connect(id []byte) error
	Pub(to []byte, data []byte) error
}