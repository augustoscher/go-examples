package main

import "fmt"

func main() {
	x := 1
	y := 2

	//Go só tem apenas posfix
	x++ // x = x + 1 | x += 1
	fmt.Println(x)

	y-- // y = y - 1 | y -= 1
	fmt.Println(y)
}
