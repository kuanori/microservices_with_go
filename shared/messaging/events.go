package messaging

import pb "microservices_with_go/shared/proto/trip"

const (
	FindAvailableDriversQueue = "find_available_drivers"
	NotifyNewTripQueue        = "notify_new_trip_queue"
)

type TripEventData struct {
	Trip *pb.Trip `json:"trip"`
}
