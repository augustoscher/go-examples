package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	s1[0] = 7 // because s1 and s2 has a pointer to the same array, the value will be displayed in both slices
	fmt.Println(s1, s2)
}
