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
	return p.name + " - " + string(p.age)
}

func (p Product) toString() string {
	return fmt.Sprintf("%s - %.2f", p.name, p.price)
}

func print(param Printable) {
	fmt.Println(param.toString())
}

func main() {

}
