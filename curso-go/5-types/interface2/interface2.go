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
}
