package main

import "fmt"

type Average float64

func (avg Average) between(firstValue, lastValue float64) bool {
	baseAvg := float64(avg)
	return baseAvg >= firstValue && baseAvg <= lastValue
}

func getResult(avg Average) string {
	if avg.between(9.0, 10.0) {
		return "A"
	} else if avg.between(6.5, 8.99) {
		return "B"
	} else if avg.between(4.0, 6.4) {
		return "C"
	}
	return "D"
}

func main() {
	fmt.Println(getResult(9.0))
	fmt.Println(getResult(7.5))
	fmt.Println(getResult(6.0))
	fmt.Println(getResult(2.25))
}
