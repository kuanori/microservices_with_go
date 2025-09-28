package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type Truck struct {
	id string
}

// method of the Truck
func (t *Truck) LoadCargo() error {
	return nil
}

func (t *Truck) UnloadCargo() error {
	return nil
}

// processTruck handles the loading and unloading of a truck
func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %s\n", truck.id)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	return nil
}

func main() {

	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrvied.\n", truck.id)

		if err := processTruck(truck); err != nil {
			log.Fatalf("Error processing truck: %s", err)
		}
	}

}
