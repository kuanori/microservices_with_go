package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	h "microservices_with_go/services/trip-service/internal/infrastructure/http"
	"microservices_with_go/services/trip-service/internal/infrastructure/repository"
	"microservices_with_go/services/trip-service/internal/service"
	"microservices_with_go/shared/env"
)

var (
	httpAddr = env.GetString("TRIP_HTTP_ADDR", ":8083")
)

func main() {
	log.Println("Trip Service started")

	// Инициализация зависимостей
	mux := http.NewServeMux()
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)
	httpHandler := h.HttpHandler{Service: svc}

	mux.HandleFunc("POST /preview", httpHandler.HandleTripPreview)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	// Канал для сигналов завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Запускаем сервер в отдельной горутине
	go func() {
		log.Printf("HTTP server listening on %s", httpAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Ожидаем сигнал остановки
	<-stop
	log.Println("Shutting down Trip Service...")

	// Контекст с таймаутом для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to gracefully shutdown HTTP server: %v", err)
	}

	log.Println("Trip Service stopped gracefully")
}
