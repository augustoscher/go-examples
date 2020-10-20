package main

import "fmt"

func printResult(grade float64) {
	if grade >= 7 {
		fmt.Println("Approved: ", grade)
	} else {
		fmt.Println("Repproved: ", grade)
	}
}

func main() {
	printResult(7.1)
	printResult(7)
	printResult(6.9)
}
