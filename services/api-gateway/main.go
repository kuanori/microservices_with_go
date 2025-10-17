package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"microservices_with_go/shared/env"
)

var (
	httpAddr = env.GetString("GATEWAY_HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /trip/preview", enableCORS(handleTripPreview))
	mux.HandleFunc("/ws/drivers", handleDriversWebsocket)
	mux.HandleFunc("/ws/riders", handleRidersWebsocket)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Server listetning on %s", httpAddr)
		serverErrors <- server.ListenAndServe()
	}()

	shutDown := make(chan os.Signal, 1)
	signal.Notify(shutDown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Printf("Error starting the server: %v", err)
	case sig := <-shutDown:
		log.Printf("Error starting the server: %v", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Could not stop server gracefully: %v", err)
			server.Close()
		}

	}
}
