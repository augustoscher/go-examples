package main

import "fmt"

func someHardProcess(n int) int {
	defer fmt.Println("Fim") // atrsar execução para o ultimo momento possível no contexto de execução do metodo.
	if n > 5000 {
		fmt.Println("High")
	} else {
		fmt.Println("Low")
	}
	return n
}

func someHardProcess2(n int) int {
	fmt.Println("Fim")
	if n > 5000 {
		fmt.Println("High")
	} else {
		fmt.Println("Low")
	}
	return n
}

func main() {
	fmt.Println(someHardProcess(5000))
	fmt.Println()
	fmt.Println(someHardProcess2(5000))
}
