package main

import "fmt"

func main() {

	microservices := make(map[string]int)
	microservices["payment"] = 1

	var input string
	fmt.Print("Поиск по названию сервиса: ")
	fmt.Scan(&input)

	if val, ok := microservices["payment"]; ok {
		fmt.Printf("Микросервис %v найден со значением %v\n", input, val)
	}

	delete(microservices, "payment")
	clear(microservices)

	_, exists := microservices["payment"]
	if !exists {
		fmt.Printf("Микросервис payment не найден\n")
	}

}
