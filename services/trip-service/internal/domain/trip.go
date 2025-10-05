package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripModel struct {
	ID       primitive.ObjectID
	UserID   string // rider id
	Status   string
	RideFare *RideFareModel
}

type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
}

type TripService interface {
	CreateTrip(ctx context.Context, fare *RideFareModel) (*TripModel, error)
}

// we will return the in-memory reference to the struct,
// instead of just having a copy
