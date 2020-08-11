package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //atribui automaticamente o tipo (float64)

	// forma reduzida - declara e inicializa
	area := PI * math.Pow(raio, 2)

	fmt.Printf("Área da circunferência de raio %f é %f\n", raio, area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	// forma reduzida
	g, h, i := false, 2, "ola"
	fmt.Println(g, h, i)

	// não permite mudar o tipo da variável
	// i = 12
}
