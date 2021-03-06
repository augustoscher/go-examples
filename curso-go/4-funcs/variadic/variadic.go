package main

import "fmt"

func sum(a, b float64) float64 {
	return a + b
}

func max(a, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}

func reduce(fn func(float64, float64) float64, numbers ...float64) float64 {
	result := 0.0
	for _, number := range numbers {
		result = fn(result, number)
	}
	return result
}

func avg(numbers ...float64) float64 {
	// total := 0.0
	// for _, number := range numbers {
	// 	total += number
	// }
	fmt.Println(reduce(max, numbers...))

	total := reduce(sum, numbers...)
	return total / float64(len(numbers))
}

func printApproved(approved ...string) {
	fmt.Println("Approved:")
	for _, value := range approved {
		fmt.Println(value)
	}
}

func main() {
	fmt.Printf("Avg: %.2f\n", avg(5.0, 7.0, 5.0))

	name := []string{
		"John",
		"Tom",
		"Hana",
	}
	// spread slice in variadic
	printApproved(name...)
}
