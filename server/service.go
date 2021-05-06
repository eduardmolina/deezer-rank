package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"net/http"
)

type Service struct {
	Post       string
	StaticPath string
	Server     *http.Server
	Redis      *redis.Client
}

func NewService(port string, staticPath string) (*Service, error) {

	handler := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	rds := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	service := &Service{port, staticPath, srv, rds}

	service.CreateUIBlueprint(handler)
	service.CreateAPIBlueprint(handler)

	return service, nil
}

func (s *Service) Start() {
	s.Server.ListenAndServe()
}

func (s *Service) Shutdown(ctx context.Context) {
	s.Server.Shutdown(ctx)
}

func (s *Service) CreateUIBlueprint(handler *http.ServeMux) {
	handler.Handle("/", http.FileServer(http.Dir(s.StaticPath)))
}

func (s *Service) CreateAPIBlueprint(handler *http.ServeMux) {
	handler.HandleFunc("/api/healthcheck", s.HealthCheck)
	handler.HandleFunc("/api/tracks", s.ListTracks)
}

func (s *Service) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alive"))
}

func (s *Service) ListTracks(w http.ResponseWriter, r *http.Request) {

}
