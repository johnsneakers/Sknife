package srv

import (
	"SKnife/conf"
	"runtime"
)

type Service interface {
	Name() string
	Version() string
	Config(conf *conf.ServiceConf) error
	Start()
}

func RunService(service Service) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	service.Start()
}