package srv

import (
	"SKnife/conf"
	"runtime"
	"fmt"
	"SKnife/utils"
)

type Service interface {
	Name() string
	Version() string
	Config(conf *conf.ServiceConf) error
	Start()
}

func RunService(service Service) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	conf, err := conf.LoadServiceConf(service.Name())
	if err != nil {
		panic(err)
	}

	extraProgName := fmt.Sprintf("-version:%s -port:%d -group:%s", service.Version(), conf.Server.Port)
	utils.SetVersion(extraProgName)
	service.Config(conf)
	service.Start()
}