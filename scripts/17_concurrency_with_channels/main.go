package main

import (
	"context"
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

type contextKey string

var UserIDKey contextKey = "userID"

// processTruck handles the loading and unloading of a truck
func processTruck(ctx context.Context, truck Truck) error {

	fmt.Printf("Started processing truck: %+v\n", truck)

	// access the user id
	hardCodeUserID := ctx.Value("hardCodeUserID")
	UserIDKey := ctx.Value(UserIDKey)
	log.Printf(" hardCodeUserID: %v, UserIDKey: %v", hardCodeUserID, UserIDKey)

	// ctx.WithTimeout + select + ctx.Done() — стандартный способ ограничить время выполнени
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel() // means - need to clean up go processes

	// simulate a long running process
	delay := time.Second * 1
	select {
	case <-ctx.Done(): // we are listening channels to be called from the context
		return ctx.Err()
	case <-time.After(delay):
		break
	}

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

func processFleet(ctx context.Context, trucks []Truck) error {
	// Создаётся WaitGroup для отслеживания горутин
	var wg sync.WaitGroup

	// channels act as pipeline for to send and receive data
	errorsChan := make(chan error, len(trucks))

	for _, t := range trucks {
		// 1. specify how many goroutins do we have
		// we add 5 goroutines
		wg.Add(1)

		go func(t Truck) { // anonymous func
			if err := processTruck(ctx, t); err != nil {
				log.Println(err)
				errorsChan <- err // Sending to chan
			}

			wg.Done()
		}(t)
	}

	// we need to wait all 5 of goroutines to call wg.Done()
	wg.Wait()
	close(errorsChan)

	var errs []error

	for err := range errorsChan {
		log.Printf("Error processing truck: %v\n", err)
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("fleet processing had %d errors", len(errs))
	}

	return nil
}

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "hardCodeUserID", 42)
	ctx = context.WithValue(ctx, UserIDKey, 228)

	trucks := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	if err := processFleet(ctx, trucks); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
		return
	}

	fmt.Print("All trucks were processed successfuly\n")
}
