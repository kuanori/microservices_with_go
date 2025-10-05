package domain

import (
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
