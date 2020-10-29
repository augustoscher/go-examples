package main

import "fmt"

// func que retorna uma outra func
func closure() func() {
	x := 10
	var fn = func() {
		fmt.Println(x)
	}
	return fn
}

func main() {
	x := 20
	fmt.Println(x)
	fn := closure()
	fn() //print x of lin 7
}
