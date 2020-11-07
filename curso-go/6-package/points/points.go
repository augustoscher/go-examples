package main

import "math"

// Visibilidade sempre a níve de pacote
// Inicializado com letra maiúscula: Publico
// Inicializado com letra maiúscula: Privado
// Quando compilado, existe apenas a estrutura de pacote

// Point represent a cordinate in cartesian plan
type Point struct {
	x float64
	y float64
}

func side(a, b Point) (sx, sy float64) {
	sx = math.Abs(b.x - a.x)
	sy = math.Abs(b.y - a.y)
	return
}

// Distance calculate the linear distance between two points
func Distance(a, b Point) float64 {
	sx, sy := side(a, b)
	return math.Sqrt(math.Pow(sx, 2) + math.Pow(sy, 2))
}
