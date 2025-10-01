package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
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

	fmt.Printf("Started processing truck: %+v\n", truck)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	// freeze for 1 sec
	time.Sleep(time.Second)

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	fmt.Printf("Finished processing truck: %+v\n", truck)
	return nil
}

func processFleet(trucks []Truck) error {
	var wg sync.WaitGroup

	for _, t := range trucks {
		// 1. specify how many goroutins do we have
		// we add 5 goroutines
		wg.Add(1)

		go func(t Truck) { // anonymous func
			if err := processTruck(t); err != nil {
				log.Println(err)
			}

			wg.Done()
		}(t)
	}

	// we need to wait all 5 of them to call wg.Done()
	wg.Wait()

	return nil
}

func main() {

	trucks := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	if err := processFleet(trucks); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
		return
	}

	fmt.Print("All trucks were processed successfuly\n")
}
