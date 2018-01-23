package api

import (
	"SKnife/srv"
	"SKnife/conf"
	"fmt"
	"net/http"
	"github.com/justinas/alice"
)

type service struct {
	conf *conf.ServiceConf
}

const (
	HELLO_VERSION = "0.0.1"
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
	fmt.Println("start " + s.conf.Name + " service listen host:", s.conf.Server.Host, " port:", s.conf.Server.Port)
	addr := fmt.Sprintf("%s:%d", s.conf.Server.Host, s.conf.Server.Port)
	common := alice.New()
	http.Handle("/hello", common.ThenFunc(s.Hello))
	if e := http.ListenAndServe(addr, nil); e != nil {
		panic(e)
	}
}
