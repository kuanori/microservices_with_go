package main

import "fmt"

type NormalTruck struct {
	id    string
	cargo int
}

func main() {

	truckID := 42
	anotherTruckID := &truckID

	fmt.Println(&truckID)       // shows the addr
	fmt.Println(anotherTruckID) // shows the addr

	truckID = 0

	fmt.Println(*anotherTruckID) // 0

	t := NormalTruck{cargo: 0}

	fillTruckCargo(t)
	fmt.Println(t) // 0

	fillTruckCargoCorrect(&t) // pass the addr!
	fmt.Println(t)            // 100

}

func fillTruckCargo(t NormalTruck) {
	t.cargo = 100
}

// func that receives pointer
func fillTruckCargoCorrect(t *NormalTruck) {
	t.cargo = 100
}
