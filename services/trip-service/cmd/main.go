package main

import (
	"log"
	h "microservices_with_go/services/trip-service/internal/infrastructure/http"
	"microservices_with_go/services/trip-service/internal/infrastructure/repository"
	"microservices_with_go/services/trip-service/internal/service"
	"microservices_with_go/shared/env"
	"net/http"
)

var (
	httpAddr = env.GetString("TRIP_HTTP_ADDR", ":8083")
)

func main() {
	log.Println("Trip Service started")

	mux := http.NewServeMux()
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)
	httpHandler := h.HttpHandler{Service: svc}

	mux.HandleFunc("POST /preview", httpHandler.HandleTripPreview)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error: %v", err)
	}
}
