package main

import "net/http"

func NewHttpServer(addr string) *http.Server {
	s := &http.Server{
		Addr: addr,
	}

	return s

}
