package main

import "fmt"

func main() {
	// arrays são homogêneos (mesmo tipo de dado dentro dele) e estática (fixo)
	var grades [3]float64
	fmt.Println(grades) // [0, 0, 0]

	grades[0], grades[1], grades[2] = 9.8, 7.4, 10.0
	fmt.Println(grades)

	total := 0.0
	for i := 0; i < len(grades); i++ {
		total += grades[i]
	}

	avg := total / float64(len(grades))
	fmt.Printf("Méida é: %.2f\n", avg)

	
}
