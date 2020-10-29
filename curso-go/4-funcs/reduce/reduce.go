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

func min(a, b float64) float64 {
	if a <= b {
		return a
	}
	return b
}

func reduce(fn func(float64, float64) float64, numbers ...float64) float64 {
	result := numbers[0]
	for _, number := range numbers {
		result = fn(result, number)
	}
	return result
}

func avg(numbers ...float64) float64 {
	total := reduce(sum, numbers...)
	return total / float64(len(numbers))
}

func main() {
	numbers := []float64{5.0, 7.5, 3.2}
	fmt.Printf("Avg: %.2f\n", avg(numbers...))
	fmt.Printf("Max: %.2f\n", reduce(max, numbers...))
	fmt.Printf("Min: %.2f\n", reduce(min, numbers...))
}
