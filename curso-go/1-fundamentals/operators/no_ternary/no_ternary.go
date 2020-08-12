package main

import "fmt"

func retrieveResults(grade float64) string {
	if grade >= 7 {
		return "Approved"
	}
	return "Reproved"
}

func main() {
	// Doesnt exists ternary operators in go
	fmt.Println(retrieveResults(7.1))
	fmt.Println(retrieveResults(5.1))
}
