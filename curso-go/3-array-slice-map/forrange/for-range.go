package main

import "fmt"

func main() {
	numbers := [...]int{9, 1, 2, 3, 4, 5} //compilador conta a qtd do array
	for i, num := range numbers {
		fmt.Printf("Position: %d - Number: %d\n", i, num)
	}

	for _, num := range numbers {
		fmt.Printf("Number: %d\n", num)
	}
}
