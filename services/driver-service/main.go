package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"microservices_with_go/shared/env"
	"microservices_with_go/shared/messaging"

	grpcserver "google.golang.org/grpc"
)

var (
	httpAddr    = env.GetString("DRIVER_HTTP_ADDR", ":9084")
	rabbitmqURI = env.GetString("RABBITMQ_URI", "amqp://guest:guest@rabbitmq:5672/")
)

var GrpcAddr = ":9084"

func main() {

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

	conn, err := messaging.NewRabbitMQ(rabbitmqURI)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Println("Starting RabbitMQ connection")

	service := NewService()

	// Starting the gRPC server
	grpcServer := grpcserver.NewServer()
	NewGrpcHandler(grpcServer, service)

	log.Printf("Starting gRPC server Driver on port %s", lis.Addr())

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
}
