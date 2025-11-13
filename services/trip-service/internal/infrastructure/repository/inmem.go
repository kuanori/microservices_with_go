package repository

import (
	"context"
	"fmt"
	"microservices_with_go/services/trip-service/internal/domain"
	pbd "microservices_with_go/shared/proto/driver"
	pb "microservices_with_go/shared/proto/trip"
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

func (r *inmemRepository) SaveRideFare(ctx context.Context, f *domain.RideFareModel) error {
	r.rideFare[f.ID.Hex()] = f

	return nil
}

func (r *inmemRepository) GetRideFareByID(ctx context.Context, id string) (*domain.RideFareModel, error) {
	fare, exist := r.rideFare[id]
	if !exist {
		return nil, fmt.Errorf("fare does not exits with ID: %v", id)
	}

	return fare, nil
}

func (r *inmemRepository) GetTripByID(ctx context.Context, id string) (*domain.TripModel, error) {
	trip, ok := r.trips[id]
	if !ok {
		return nil, nil
	}
	return trip, nil
}

func (r *inmemRepository) UpdateTrip(ctx context.Context, tripID string, status string, driver *pbd.Driver) error {
	trip, ok := r.trips[tripID]
	if !ok {
		return fmt.Errorf("trip not found with ID: %s", tripID)
	}

	trip.Status = status

	if driver != nil {
		trip.Driver = &pb.TripDriver{
			Id:             driver.Id,
			Name:           driver.Name,
			CarPlate:       driver.CarPlate,
			ProfilePicture: driver.ProfilePicture,
		}
	}
	return nil
}
