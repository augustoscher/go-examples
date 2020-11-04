package main

import "fmt"

type Printable interface {
	toString() string
}

type Person struct {
	name string
	age  int
}

type Product struct {
	name  string
	price float64
}

// interfaces são implementadas implicitamente
func (p Person) toString() string {
	return fmt.Sprintf("%s - %d", p.name, p.age)
}

func (p Product) toString() string {
	return fmt.Sprintf("%s - %.2f", p.name, p.price)
}

func print(param Printable) {
	fmt.Println(param.toString())
}

func main() {
	var person Printable = Person{
		name: "Augusto",
		age:  29,
	}

	var product Printable = Product{
		name:  "Palito",
		price: 10.00,
	}

	print(person)
	print(product)
	print(Product{"P2", 149.99})
	fmt.Println()
	fmt.Println(person.toString())
	fmt.Println(product.toString())
}
