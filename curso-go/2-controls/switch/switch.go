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

func getFromTime() {
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

func getValueFromType(i interface{}) string {
	switch i.(type) {
	case int:
		return "integer"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "unknow"
	}
}

func main() {
	fmt.Println(gradeToConcept(10))
	fmt.Println(gradeToConcept(5))

	getFromTime()

	fmt.Println(getValueFromType(2.9))
	fmt.Println(getValueFromType(1))
	fmt.Println(getValueFromType("oi"))
	fmt.Println(getValueFromType(func() {}))
	fmt.Println(getValueFromType(time.Now()))
}
