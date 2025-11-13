package messaging

import pb "microservices_with_go/shared/proto/trip"

const (
	FindAvailableDriversQueue = "find_available_drivers"
	NotifyNewTripQueue        = "notify_new_trip_queue"
	DriverCmdTripRequestQueue = "driver_cmd_trip_request"
)

type TripEventData struct {
	Trip *pb.Trip `json:"trip"`
}
