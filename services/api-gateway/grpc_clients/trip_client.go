package grpc_clients

import (
	pb "microservices_with_go/shared/proto/trip"
	"microservices_with_go/shared/tracing"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type tripServiceClient struct {
	Client pb.TripServiceClient // from shared/proto/trip/trip_grpc.pb.go
	conn   *grpc.ClientConn
}

func NewTripServiceClient() (*tripServiceClient, error) {
	tripServiceURL := os.Getenv("TRIP_SERVICE_URL")
	if tripServiceURL == "" {
		tripServiceURL = "trip-service:9083" // from infra/development/k8s/trip-service-deployment.yaml
	}

	dialOptions := append(
		tracing.DialOptionsWithTracing(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	conn, err := grpc.NewClient(tripServiceURL, dialOptions...)
	if err != nil {
		return nil, err
	}

	client := pb.NewTripServiceClient(conn) // from shared/proto/trip/trip_grpc.pb.go

	return &tripServiceClient{
		Client: client,
		conn:   conn,
	}, nil

}

func (c *tripServiceClient) Close() {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return
		}
	}
}
