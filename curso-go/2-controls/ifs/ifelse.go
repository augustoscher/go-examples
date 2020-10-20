package main

import "fmt"

func printResult(grade float64) string {
	if grade >= 7 {
		return "Approved"
	} else {
		return "Repproved"
	}
}

func gradeToConcept(grade float64) string {
	if grade >= 9 && grade <= 10 {
		return "A"
	} else if grade >= 8 && grade < 9 {
		return "B"
	} else if grade >= 5 && grade < 8 {
		return "C"
	} else if grade < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	grade := 9.0
	fmt.Printf("Grade: %f - Status: %s - Concept: %s\n", grade, printResult(grade), gradeToConcept(grade))

	grade = 8
	fmt.Printf("Grade: %f - Status: %s - Concept: %s\n", grade, printResult(grade), gradeToConcept(grade))

	grade = 5
	fmt.Printf("Grade: %f - Status: %s - Concept: %s\n", grade, printResult(grade), gradeToConcept(grade))

	grade = 4
	fmt.Printf("Grade: %f - Status: %s - Concept: %s\n", grade, printResult(grade), gradeToConcept(grade))

	grade = 11
	fmt.Printf("Grade: %f - Status: %s - Concept: %s\n", grade, printResult(grade), gradeToConcept(grade))
}
