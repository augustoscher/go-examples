package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6) error: first argument need to be a slice
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1) //[1 2 3] [4 5 6]

	slice2 := make([]int, 2)
	copy(slice2, slice1) //slice2 is length 2. It will have [4, 5]
	fmt.Println(slice1, slice2)
}
