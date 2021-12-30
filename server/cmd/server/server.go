package main

import (
	"context"
	"fmt"
	"net/http"
)

type HttpPortNumber int

type ApiServer struct {
	Port   HttpPortNumber
	router *Router
	server *http.Server
}

func (s *ApiServer) Start() error {

	handler := new(http.ServeMux)
	handler.HandleFunc("/", s.router.handle)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

func (s *ApiServer) Stop() error {
	if s.server != nil {
		return s.server.Shutdown(context.Background())
	} else {
		return fmt.Errorf("Server did not start")
	}
}
