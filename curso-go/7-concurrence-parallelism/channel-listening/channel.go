package main

import (
	"fmt"
	"time"
)

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
	// encerra a comunicação do canal
	// sem o close, teremos um deadlock no range do channel,
	// pois ele estará lendo dados mas nenhuma go routine estará escrevendo.
	close(c)
}

func main() {
	c := make(chan int, 30)
	go getPrimos(cap(c), c) // func cap retorna o tamanho do buffer do channel
	i := 0
	for numeroPrimo := range c {
		i++
		fmt.Printf("Primo %d: %d\n", i, numeroPrimo)
	}
	fmt.Println("Fim")
}
