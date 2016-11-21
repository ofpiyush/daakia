package daakia


type Method struct {
	Name string
	In bool
}

type Service struct {
	Name string
	Namespace string
	Server map[byte]*Method
	Client map[byte]*Method
}

func NewService(name, namespace string) *Service {
	return &Service{
		Name: name,
		Namespace: namespace,
		Server: make(map[byte]*Method),
		Client: make(map[byte]*Method),
	}
}

