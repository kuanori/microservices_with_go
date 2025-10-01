package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		t.Run("should loand and unload a truck cargo", func(t *testing.T) {
			nt := NormalTruck{id: "1", cargo: 0}
			et := ElectricTruck{id: "2"}

			if err := processTruck(&nt); err != nil {
				t.Fatalf("Error processing truck: %s", err)
			}

			if err := processTruck(&et); err != nil {
				t.Fatalf("Error processing truck: %s", err)
			}

			if nt.cargo != 0 {
				t.Fatalf("Normal truck cargo should be 0")
			}

			if et.battery > 0 {
				t.Fatalf("Electric truck battery should less than 0")
			}
		})
	})
}
