package main

import "fmt"

type Course struct {
	name string
}

func main() {
	var something interface{}
	fmt.Println(something)

	something = 3
	fmt.Println(something)

	type Dynamic interface{}
	var something2 Dynamic = "Opa"
	fmt.Println(something2)
}
