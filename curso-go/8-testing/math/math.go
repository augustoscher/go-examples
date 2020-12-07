package math

//Average is responsible for calculate the sum of two numbers divided by 2
func Average(numbers ...float64) float64 {
	sum := 0.0
	for _, n := range numbers {
		sum += n
	}
	return sum / float64(len(numbers))
}
