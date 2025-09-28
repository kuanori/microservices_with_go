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

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

// so the NormalTruck now implements Truck interface
func (t *NormalTruck) LoadCargo() error {
	t.cargo += 1
	return nil
}

// когда добавляем указатель, то мы работаем с оригиналом, а не копией, и в вызывающей функции должны передавать через &
func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

// so the ElectricTruck now implements Truck interface
func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery -= 1
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery -= 1
	return nil
}

// processTruck handles the loading and unloading of a truck
func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %+v\n", truck)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}

	return nil
}

func main() {
	nt := NormalTruck{id: "1"}
	et := ElectricTruck{id: "2"}

	person := make(map[string]any, 0)
	person["name"] = "Kuan"
	person["age"] = 24

	_, exists := person["age"]
	if !exists {
		log.Fatal("age not exists")
		return
	}

	if err := processTruck(&nt); err != nil {
		log.Fatalf("Error processing truck: %s", err)
	}

	if err := processTruck(&et); err != nil {
		log.Fatalf("Error processing truck: %s", err)
	}

	log.Println(nt)
	log.Println(et)

}
