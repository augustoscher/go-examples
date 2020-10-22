package main

import (
	"fmt"
	"reflect"
)

// slice: Representa um pedaço de um array
// flexibiliza o uso do array
// array tem tamanho fixo definido
// slice tem um tamanho variável

func main() {
	array1 := [3]int{1, 2, 3}
	slice1 := []int{1, 2, 3}

	fmt.Println(array1, slice1)
	fmt.Println(reflect.TypeOf(array1), reflect.TypeOf(slice1))

	array2 := [5]int{5, 4, 3, 2, 1}
	slice2 := array2[1:3]
	fmt.Println(slice2) // 4, 3 pois o ultimo elemento não é incluído.

	s3 := array2[:3]
	fmt.Println(s3) //5, 4, 3

	s4 := array2[3:]
	fmt.Println(s4) //2, 1

	sliceOfSlice := s3[:1]
	fmt.Println(sliceOfSlice) //5

}
