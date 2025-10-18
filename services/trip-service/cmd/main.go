package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"microservices_with_go/services/trip-service/internal/infrastructure/repository"
	"microservices_with_go/services/trip-service/internal/service"
	"microservices_with_go/shared/env"

	grpcserver "google.golang.org/grpc"
)

var (
	httpAddr = env.GetString("TRIP_HTTP_ADDR", ":8083")
)

var GrpcAddr = ":9093"

func main() {
	mux := http.NewServeMux()
	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		cancel()
	}()

	// ======== GRPc Server
	lis, err := net.Listen("tcp", GrpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpcserver.NewServer()
	// TODO initialize our grpc handler implementation

	log.Printf("Starting gRPC server Trip on port %s", lis.Addr())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to server: %v", err)
			cancel()
		}
	}()

	// wait for the shutdown signal
	<-ctx.Done()
	log.Println("Shutting down the server...")
	grpcServer.GracefulStop()

	// ======== HTTP Server
	// httpHandler := h.HttpHandler{Service: svc}

	// mux.HandleFunc("POST /preview", httpHandler.HandleTripPreview)

	// server := &http.Server{
	// 	Addr:    httpAddr,
	// 	Handler: mux,
	// }

	// Канал для сигналов завершения
	// stop := make(chan os.Signal, 1)
	// signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Запускаем сервер в отдельной горутине
	// go func() {
	// 	log.Printf("HTTP server listening on %s", httpAddr)
	// 	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("HTTP server error: %v", err)
	// 	}
	// }()

	// Ожидаем сигнал остановки
	// <-stop
	// log.Println("Shutting down Trip Service...")

	// // Контекст с таймаутом для graceful shutdown
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// if err := server.Shutdown(ctx); err != nil {
	// 	log.Fatalf("Failed to gracefully shutdown HTTP server: %v", err)
	// }

	// log.Println("Trip Service stopped gracefully")
}
