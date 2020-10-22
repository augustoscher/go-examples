package main

import "fmt"

// slice: flexibiliza o uso do array
// array tem tamanho fixo definido
// slice tem um tamanho variável

func main() {
	array1 := [3]int{1, 2, 3}
	slice1 := []int{1, 2, 3}

	fmt.Println(array1, slice1)
}
