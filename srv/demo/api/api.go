package api

import (
	"net/http"
	"fmt"
)

func (s *service) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helloWorld!")

}