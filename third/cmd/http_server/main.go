package main

import (
	"log"
	"net/http"

	"github.com/KasiditR/backend-challenge-third/internal/handlers"
)

const (
	port = ":8080"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/beef", handlers.GetBeefHandler)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Printf("HTTP server is running on port %v", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
