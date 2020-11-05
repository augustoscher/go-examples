package main

import "fmt"

type Runnable interface {
	run()
}

type Animal struct {
	name    string
	age     int
	running bool
}

func (a *Animal) run() {
	a.running = true
}

func main() {
	pernaLonga := Animal{"PernaLonga", 20, false}
	fmt.Println(pernaLonga)
	pernaLonga.run()
	fmt.Println(pernaLonga)

	// Passa a referência do Animal criado para a variavel patolino
	var patolino = &Animal{"Patolino", 22, false}
	fmt.Println(patolino)
	patolino.run()
	fmt.Println(patolino)
}
