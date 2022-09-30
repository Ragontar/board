package server

import (
	"net/http"
)

var serverAddr string

func Init() {
	serverAddr = "0.0.0.0:9010"
}

func Run() error {
	router := NewRouter()
	return http.ListenAndServe(serverAddr, router)
}
