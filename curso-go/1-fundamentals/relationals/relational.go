package main

import (
	"fmt"
	"time"
)

func main() {
	// relational operators returns bool value
	fmt.Println("Strings compare:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)   // false
	fmt.Println(">", 3 > 2)   // true
	fmt.Println(">=", 3 >= 2) // true
	fmt.Println("<=", 3 <= 2) // false

	//dates
	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	// ambos não comparam usando endereço de memória e sim o dado
	fmt.Println("Dates:", d1 == d2)
	fmt.Println("Dates:", d1.Equal(d2))

	// struct
	type Pessoa struct {
		Name string
	}

	// também não olha para a referência de memoria
	p1 := Pessoa{"Joao"}
	p2 := Pessoa{"Maria"}
	p3 := Pessoa{"Maria"}
	fmt.Println("Structs:", p1 == p2) // false
	fmt.Println("Structs:", p2 == p3) // true

}
