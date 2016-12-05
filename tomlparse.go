package daakia

import (
	"os"

	"github.com/BurntSushi/toml"
)

type ServiceConf struct {
	Name   string
	Server []*Method
	Client []*Method
}

type TomlConf struct {
	Namespace string
	Services  []*ServiceConf
}

func ParseToml(file *os.File, skip int) ([]*Service, error) {
	var conf TomlConf
	_, err := toml.DecodeFile(file.Name(), &conf)
	if err != nil {
		return nil, err
	}

	services := make([]*Service, len(conf.Services))
	for i, service := range conf.Services {
		services[i] = NewService(service.Name, conf.Namespace)
		curSkip := uint32(skip)
		for _, serverMethod := range service.Server {
			services[i].Server[byte(curSkip)] = serverMethod
			curSkip++
		}
		for _, clientMethod := range service.Client {
			services[i].Client[byte(curSkip)] = clientMethod
			curSkip++
		}
	}
	return services, nil
}
