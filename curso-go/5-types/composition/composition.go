package main

import "fmt"

type Animal struct {
	name   string
	weight int
}

type Cow struct {
	Animal   //Campos anonimos. Composição e não herança
	pregnant bool
}

func main() {
	c := Cow{}
	c.name = "Mimosa"
	c.weight = 350
	c.pregnant = true
	fmt.Println(c)
	fmt.Printf("Cow: %s - Weight: %d - Pregnant: %v\n", c.name, c.weight, c.pregnant)

}
