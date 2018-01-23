package api

import (
	"SKnife/srv"
	"SKnife/conf"
	"fmt"
)

type service struct {
	name string
}

const (
	HELLO_VERSION = "1.9.4"
)

func NewService() srv.Service {
	return new(service)
}

func (s *service) Name() string {
	return "demoService!"
}

func (s *service) Version() string {
	return HELLO_VERSION
}

func (s *service) Config(conf *conf.ServiceConf) error {
	return nil
}

func (s *service) Start() {
	// start from here...
	fmt.Println("start from here....")
}