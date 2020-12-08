package mathematics

import (
	"fmt"
	"strconv"
)

//Average is responsible for calculate the sum of two numbers divided by 2
func Average(numbers ...float64) float64 {
	sum := 0.0
	for _, n := range numbers {
		sum += n
	}
	avg := sum / float64(len(numbers))
	rounded, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", avg), 64)
	return rounded
}
