package conf
type ServerConf struct {
	Host         string
	Port         int
}


type ServiceConf struct {
	Path     string
	Name     string
	Mode     string
	Server   *ServerConf
}