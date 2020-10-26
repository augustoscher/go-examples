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

	//loop in map
	for key, value := range approved {
		fmt.Printf("ID: %d - Name: %s\n", key, value)
	}

	//read from map
	fmt.Printf("Map value: %s\n", approved[298472983])

	//delete value
	delete(approved, 987654321)

	fmt.Println(approved) //tom and john
}
