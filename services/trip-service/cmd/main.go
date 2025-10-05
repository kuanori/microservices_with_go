package main

import (
	"context"
	"log"
	"microservices_with_go/services/trip-service/internal/domain"
	"microservices_with_go/services/trip-service/internal/infrastructure/repository"
	"microservices_with_go/services/trip-service/internal/service"
	"time"
)

func main() {

	ctx := context.Background()

	inmemRepo := repository.NewInmemRepository()

	svc := service.NewService(inmemRepo)

	fare := &domain.RideFareModel{
		UserID: "42",
	}

	t, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}

	log.Println(t)

	// keep the program running for now
	for {
		time.Sleep(time.Second)
	}
}
