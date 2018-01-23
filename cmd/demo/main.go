package main

import (
	"SKnife/srv"
	"SKnife/srv/demo/api"
)

func main() {
	srv.RunService(api.NewService())
}
