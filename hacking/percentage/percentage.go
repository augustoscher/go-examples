package main

import "fmt"

func percent(value, target float64) float64 {
	mult := value * float64(100)
	return mult / target
}

func main() {
	var value, target float64 = 23, 93
	fmt.Printf("Percent %f of %f is %f\n", value, target, percent(value, target))
}
