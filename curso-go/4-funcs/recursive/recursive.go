package main

import "fmt"

func factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Invalid Number %d", n)
	case n == 0:
		return 1, nil
	default:
		lastResult, _ := factorial(n - 1)
		return n * lastResult, nil
	}
}

func main() {
	res, _ := factorial(5)
	fmt.Printf("Fac: %d\n", res)

	res, err := factorial(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
