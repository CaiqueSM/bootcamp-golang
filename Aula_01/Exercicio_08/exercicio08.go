package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

func main() {
	count := 0

	fmt.Println("A idade de Benjamin é:", employees["Benjamin"])

	for _, value := range employees {
		if value > 21 {
			count++
		}
	}

	fmt.Println("O número de funcionários com mais de 21 anos é:", count)

	employees["Federico"] = 25

	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)
}
