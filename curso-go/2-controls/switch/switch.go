package main

import (
	"fmt"
	"time"
)

func gradeToConcept(grade float64) string {
	g := int(grade)
	switch g {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 3, 4:
		return "D"
	case 2, 1, 0:
		return "D"
	default:
		return "Invalid"
	}
}

func main() {
	fmt.Println(gradeToConcept(10))
	fmt.Println(gradeToConcept(5))

	t := time.Now()
	switch {
	// entra na primeira opção que der verdadeiro
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}

}
