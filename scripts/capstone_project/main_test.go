package main

import (
	"testing"
)

func TestAddTruck(t *testing.T) {
	manager := NewTruckManager()
	manager.AddTruck("1", 100)

	if len(manager.trucks) != 1 {
		t.Errorf("Expected 1 truck, got %d", len(manager.trucks))
	}
}

func TestGetTruck(t *testing.T) {
	manager := NewTruckManager()
	manager.AddTruck("1", 100)

	truck, err := manager.GetTruck("1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if truck.ID != "1" {
		t.Errorf("Expected truck ID to be 1, got %s", truck.ID)
	}
}

func TestRemoveTruck(t *testing.T) {
	manager := NewTruckManager()
	manager.AddTruck("1", 100)

	manager.RemoveTruck("1")

	_, err := manager.GetTruck("1")
	if err != ErrTruckNotFound {
		t.Errorf("Expected truck not found error, got %v", err)
	}

	if len(manager.trucks) != 0 {
		t.Errorf("Expected 0 trucks, got %d", len(manager.trucks))
	}
}

func TestUpdateTruckCargo(t *testing.T) {
	manager := NewTruckManager()
	manager.AddTruck("1", 100)

	manager.UpdateTruckCargo("1", 200)

	truck, err := manager.GetTruck("1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if truck.Cargo != 200 {
		t.Errorf("Expected truck cargo to be 200, got %d", truck.Cargo)
	}
}

/*
FIXME: Uncomment me for the concurrency part
func TestConcurrentUpdate(t *testing.T) {
	manager := NewTruckManager()
	manager.AddTruck("1", 100)
	const numGoroutines = 100
	const iterations = 100
	done := make(chan bool)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				truck, _ := manager.GetTruck("1")
				manager.UpdateTruckCargo("1", truck.Cargo+1)
			}
			done <- true
		}()
	}
	// Wait for all goroutines to complete
	for i := 0; i < numGoroutines; i++ {
		<-done
	}
} */
