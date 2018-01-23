package api

import (
	"SKnife/srv"
	"SKnife/conf"
	"fmt"
)

type service struct {
	conf *conf.ServiceConf
}

const (
	HELLO_VERSION = "1.9.4"
)

func NewService() srv.Service {
	return new(service)
}

func (s *service) Name() string {
	return "demoService"
}

func (s *service) Version() string {
	return HELLO_VERSION
}

func (s *service) Config(conf *conf.ServiceConf) error {
	s.conf = conf
	return nil
}

func (s *service) Start() {
	// start code from here:
	fmt.Println("start from here....service name:", s.conf.Name)
}