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
	// sequencial
	// speak("Joao", "Oi, tudo bem?", 3)
	// speak("Maria", "Oi, tudo certo e vc?", 1)

	//goroutines
	go speak("Joao", "Hey..", 500)
	go speak("Maria", "Ohh", 550)

	//A thread principal deve estar funcionando para as
	//goroutines executarem.
	time.Sleep(time.Second * 100)

	// Caso fizessemos assim:
	// go speak("Joao", "Hey..", 10) //goroutine
	// speak("Maria", "Ohh", 5)      //thread principal
	// Não daria tempo suficiente para Joao "falar todas as vezes"

}
