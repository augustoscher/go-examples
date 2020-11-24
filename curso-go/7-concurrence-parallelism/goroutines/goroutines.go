package main

import (
	"fmt"
	"time"
)

func speak(person, text string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%d - %s: %s\n", i+1, person, text)
	}
}

func main() {
	speak("Augusto", "Oi", 3)
}
