package http

import (
	"encoding/json"
	"log"
	"microservices_with_go/services/trip-service/internal/domain"
	"net/http"
)

type HttpHandler struct {
	Service domain.TripService
}

func (s *HttpHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	fare := &domain.RideFareModel{
		UserID: "42",
	}

	t, err := s.Service.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}

	log.Println(t)

	writeJson(w, http.StatusOK, t)
}

func writeJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
