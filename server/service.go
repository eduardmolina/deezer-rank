package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"io"
	"log"
	"net/http"
	"time"
)

type Service struct {
	Port       string
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Alive"))
}

func RefreshTracks() (string, error) {

	resp, err := http.Get("https://api.deezer.com/chart/0/tracks?limit=50")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (s *Service) ListTracks(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	deezerData, err := s.Redis.Get(ctx, "tracks").Result()
	if err != nil {
		log.Println("[+] Refreshing tracks")
		deezerData, err := RefreshTracks()
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to refresh tracks", http.StatusInternalServerError)
			return
		}

		s.Redis.Set(ctx, "tracks", deezerData, 10*time.Minute)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(deezerData))
}
