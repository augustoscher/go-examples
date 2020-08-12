package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum =>", a+b)
	fmt.Println("Sub =>", a-b)
	fmt.Println("Div =>", a/b)
	fmt.Println("Mult =>", a*b)
	fmt.Println("Mod =>", a%b)
	fmt.Println("Even =>", a%2 == 0)

	//bitwise
	fmt.Println("E =>", a&b)   //11 & 10 = 10 res = 2
	fmt.Println("Ou =>", a|b)  //11 | 10 = 11 res = 3
	fmt.Println("Xor =>", a^b) //11 ^ 10 = 01 res = 1

	// math
	c, d := 3.0, 2.0
	fmt.Println("Max valor =>", math.Max(float64(a), float64(b))) // precisa converter p float64
	fmt.Println("Min valor =>", math.Min(c, d))                   // precisa converter p float64
	fmt.Println("Pow =>", math.Pow(c, d))                         //9
}
