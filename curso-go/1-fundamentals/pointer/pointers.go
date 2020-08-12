package main

import "fmt"

func changeOnRef(pointer *int) {
	*pointer = 600
}

func main() {
	i := 1 //int - 8 bytes | 64 bits

	var p *int = nil // creating pointer
	p = &i           // get memory address of i and set in pointer
	*p++             // get value from memory address and increment
	fmt.Println("Pointer:", p)
	fmt.Println("Variable:", i)

	i++                                  // update variable.
	fmt.Println("Pointer:", p)           // same address
	fmt.Println("Variable address:", &i) // same address
	fmt.Println("Variable:", i)          // 3
	fmt.Println("Unreferencing:", *p)    // 3

	fmt.Println()
	x := 10
	px := &x
	fmt.Println("X Before:", x)
	changeOnRef(px)
	fmt.Println("X After:", x)
}
