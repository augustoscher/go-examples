package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func randomNum() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

// initializing variable before

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

	// initialize in initilizing block; Can be done in switcth, for...
	if i := randomNum(); i > 5 {
		fmt.Printf("Number %d is greather than 5.\n", i)
	} else {
		fmt.Printf("Number %d is lower than 5.\n", i)
	}

	//fmt.Println(i) i is avaiable only in if else block
}
