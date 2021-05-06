package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	service, err := NewService("3000", "./static")
	if err != nil {
		log.Fatal(err)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go service.Start()

	log.Printf("[+] Running on port %v", service.Port)
	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	service.Server.Shutdown(ctx)

	log.Printf("[+] Graceful shutdown completed")
}
