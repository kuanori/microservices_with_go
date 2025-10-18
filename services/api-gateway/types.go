package main

import (
	pb "microservices_with_go/shared/proto/trip"
	"microservices_with_go/shared/types"
)

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (p *previewTripRequest) ToProto() *pb.PreviewTripRequest {
	return &pb.PreviewTripRequest{
		UserID: p.UserID,
		StartLocation: &pb.Coordinate{
			Latitude:   p.Pickup.Latitude,
			Longtitude: p.Pickup.Longitude,
		},
		EndLocation: &pb.Coordinate{
			Latitude:   p.Destination.Latitude,
			Longtitude: p.Destination.Longitude,
		},
	}
}
