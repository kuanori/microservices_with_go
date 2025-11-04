package main

import (
	pb "microservices_with_go/shared/proto/driver"
)

type Service struct {
	drivers []*driverInMap
}

type driverInMap struct {
	Driver *pb.Driver
}

func NewService() *Service {
	return &Service{
		drivers: make([]*driverInMap, 0),
	}
}

// TODO: Register and Unregister methods
