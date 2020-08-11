package main

import "fmt"

func main() {
	// Quando não inicializamos a variável com o valor,
	// precisamos usar essa construção, definindo o tipo.

	var a int
	var b float64
	var c bool
	var d string
	var e *int //integer pointer

	fmt.Printf("%v %v %v %v %v\n", a, b, c, d, e) // 0 0 false  <nil>
	fmt.Printf("%v %v %v %q %v\n", a, b, c, d, e) // 0 0 false "" <nil>
}
