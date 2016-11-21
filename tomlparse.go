package daakia

import (
	"os"
	"github.com/BurntSushi/toml"
)


type ServiceConf struct {
	Name string
	Server []*Method
	Client []*Method
}

type TomlConf struct {
	Namespace string
	Services []*ServiceConf
}

func ParseToml(file *os.File,skip int) ([]*Service, error) {
	var conf TomlConf
	_, err := toml.DecodeFile(file.Name(),&conf)
	if err !=nil {
		return nil,err
	}

	services := make([]*Service,len(conf.Services))
	for i, service := range conf.Services {
		services[i] = NewService(service.Name,conf.Namespace)
		cur_skip := uint32(skip)
		for _, server_method :=  range service.Server {

			services[i].Server[byte(cur_skip)] = server_method
			cur_skip++
		}
		for _, client_method :=  range service.Client {
			services[i].Client[byte(cur_skip)] = client_method
			cur_skip++
		}
	}
	return services,nil
}
