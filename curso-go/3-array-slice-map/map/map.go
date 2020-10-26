package main

import "fmt"

func main() {
	// var approved map[int]string maps shoul be initialized
	//map witk int key and string value

	approved := make(map[int]string)
	approved[123456789] = "John"
	approved[987654321] = "Mary"
	approved[298472983] = "Tom"

	fmt.Println(approved)

	//override value in map
	approved[987654321] = "Mary2"

	//loop in map
	for key, value := range approved {
		fmt.Printf("ID: %d - Name: %s\n", key, value)
	}

	//read from map
	fmt.Printf("Map value: %s\n", approved[298472983])

	//delete value
	delete(approved, 987654321)

	fmt.Println(approved) //tom and john
	fmt.Println()
	// ==================================================

	//initilizing map
	salary := map[string]float64{
		"Joseph Slater": 11325.45,
		"Tom Silver":    12900.79,
		"August Hawk":   10000.00,
	}

	fmt.Println(salary)

	//writing on map
	salary["Richard Noris"] = 9500.00
	fmt.Println(salary)

	delete(salary, "NonExisting")

	for name, salary := range salary {
		fmt.Printf("Name: %s - Salary: %.2f\n", name, salary)
	}
	fmt.Println()
	// ==================================================

	//map aninhado
	salaryByLetter := map[string]map[string]float64{
		"J": {
			"Joseph Slater": 11325.45,
			"Jimmy Hendrix": 90000.00,
		},
		"T": {
			"Tom Silver": 12900.79,
		},
		"A": {
			"August Hawk": 10000.00,
		},
	}

	fmt.Println(salaryByLetter)

	//salary of jimmy hendrix
	fmt.Println(salaryByLetter["J"]["Jimmy Hendrix"])

	for name, salary := range salaryByLetter["J"] {
		fmt.Printf("Name: %s - Salary: %.2f\n", name, salary)
	}

	delete(salaryByLetter, "T")

}
