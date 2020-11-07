package main

// Running:
// cd points && go run *.go

import "fmt"

func main() {
	p1 := Point{2.0, 2.0}
	p2 := Point{2.0, 4.0}

	// side func is visible because is the same package
	fmt.Println(side(p1, p2))

	//Distance is also visible inside and outside of the package
	fmt.Println(Distance(p1, p2))
}
