package main

import "fmt"

type Run interface {
	run()
}

type Swim interface {
	swim()
}

type TriAthlete interface {
	Run
	Swim
	getName() string
}

type Athlete struct {
	name string
}

func (a Athlete) swim() {
	fmt.Println("Swimming")
}

func (a Athlete) run() {
	fmt.Println("Running")
}

func (a Athlete) getName() string {
	return a.name
}

func main() {
	a := Athlete{"Augusto"}
	fmt.Println(a)
	a.run()
	a.swim()
}
