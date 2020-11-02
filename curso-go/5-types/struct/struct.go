package main

import "fmt"

type Product struct {
	name  string
	price float64
	qtd   float64
}

// Método: func com receiver
func (p Product) total() float64 {
	return p.qtd * p.price
}

func main() {
	var p1 Product
	p1 = Product{
		name:  "P1",
		price: 1499.99,
		qtd:   5,
	}

	fmt.Println("P1: ", p1)
	fmt.Println("P1 Total: ", p1.total())
}
