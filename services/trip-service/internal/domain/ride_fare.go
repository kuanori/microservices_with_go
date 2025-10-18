package domain

import (
	pb "microservices_with_go/shared/proto/trip"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RideFareModel struct {
	ID                primitive.ObjectID
	UserID            string // rider id
	PackageSlug       string // ex: van, luxury, sedan
	TotalPriceInCents float64
	Expires           time.Time
}

func (r *RideFareModel) ToProto() *pb.RideFare {
	return &pb.RideFare{
		ID:                r.ID.Hex(),
		UserID:            r.UserID,
		PackageSlug:       r.PackageSlug,
		TotalPriceInCents: r.TotalPriceInCents,
	}
}

func ToRideFaresProto(fares []*RideFareModel) []*pb.RideFare {
	var protoFares []*pb.RideFare

	for _, f := range fares {
		protoFares = append(protoFares, f.ToProto())
	}

	return protoFares
}
