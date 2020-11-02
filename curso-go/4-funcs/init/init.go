package main

import "fmt"

// init method is executed before main
func init() {
	fmt.Println("Init...")
}

func main() {
	fmt.Println("Main...")
}

// output
// Init...
// Main...
