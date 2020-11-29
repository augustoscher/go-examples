package main

import "time"

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func getPrimos(qtd int, c chan int) {
	init := 2
	for i := 0; i < qtd; i++ {
		for val := init; ; val++ {
			if isPrimo(val) {
				c <- val
				init = val + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
}

func main() {

}
