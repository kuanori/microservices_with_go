package repository

import (
	"context"
	"microservices_with_go/services/trip-service/internal/domain"
)

type inmemRepository struct {
	trips    map[string]*domain.TripModel
	rideFare map[string]*domain.RideFareModel
}

func NewInmemRepository() *inmemRepository {
	return &inmemRepository{
		trips:    make(map[string]*domain.TripModel),
		rideFare: make(map[string]*domain.RideFareModel),
	}
}

func (r *inmemRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	r.trips[trip.ID.Hex()] = trip
	return trip, nil
}
