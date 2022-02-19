package main

import (
	"net/http"
)

type httpServer struct {
	port    string
	address string
}

func newServer(address string, port string) *httpServer {
	r := httpServer{
		address: address,
		port:    port,
	}
	return &r
}

func (t httpServer) serve() {
	address := t.address + ":" + t.port
	http.Handle("/", httpHandleRoot())
	http.ListenAndServe(address, nil)
}
