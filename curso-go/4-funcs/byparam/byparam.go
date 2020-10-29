package main

import "fmt"

func mult(a, b int) int {
	return a * b
}

// map, reduce, filter are use cases for using this
func exec(fn func(int, int) int, p1, p2 int) int {
	return fn(p1, p2)
}

func main() {
	fmt.Println(exec(mult, 5, 5))
}
