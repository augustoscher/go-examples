package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name     string
	lastname string
}

func (p Person) getCompleteName() string {
	return p.name + " " + p.lastname
}

//receiver is a pointer to object
func (p *Person) setName(name string) {
	p.name = name
}

func test(s string) {
	each := strings.Split(s, " ")
	fmt.Println("S1: ", each[0])
	fmt.Println("S2: ", each[1])
}

func main() {
	p := Person{
		name:     "Augusto",
		lastname: "Scher",
	}

	fmt.Println("Person: ", p.getCompleteName())
	p.setName("Daniela")
	fmt.Println("Person: ", p.getCompleteName())

	test("AAAAAA BBBBBB")
}
