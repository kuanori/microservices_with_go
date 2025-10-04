package main

import (
	"errors"
	"fmt"
)

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (*Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

func (m *truckManager) AddTruck(id string, cargo int) error {
	if _, exists := m.trucks[id]; exists {
		return fmt.Errorf("грузовик с %v уже существует", id)
	}

	m.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (m *truckManager) GetTruck(id string) (*Truck, error) {
	truck, exists := m.trucks[id]

	if !exists {
		return nil, ErrTruckNotFound
	}

	return truck, nil
}

func (m *truckManager) RemoveTruck(id string) error {
	_, exists := m.trucks[id]
	if !exists {
		return ErrTruckNotFound
	}

	delete(m.trucks, id)

	return nil
}

func (m *truckManager) UpdateTruckCargo(id string, cargo int) error {
	truck, exists := m.trucks[id]
	if !exists {
		return ErrTruckNotFound
	}

	truck.Cargo = cargo
	m.trucks[id] = truck

	return nil
}

type truckManager struct {
	trucks map[string]*Truck
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}
