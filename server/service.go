package main

import (
	"net/http"
)

type Service struct {
	Port       string
	StaticPath string
	Server     *http.Server
}

func NewService(port string, staticPath string) (*Service, error) {

	handler := http.NewServeMux()
	handler.Handle("/", http.FileServer(http.Dir(staticPath)))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	return &Service{port, staticPath, srv}, nil
}

func (s *Service) Start() {
	s.Server.ListenAndServe()
}
