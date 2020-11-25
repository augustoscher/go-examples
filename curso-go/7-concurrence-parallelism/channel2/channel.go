package main

import (
	"fmt"
	"time"
)

// Channel é um ponto de sincronismo.
// É a forma de comunicação entre goroutines.

func times(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base //escrevendo no canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	ch := make(chan int)
	go times(2, ch) // executa de forma independete
	fmt.Println("Não espera a goroutine")

	a, b := <-ch, <-ch // lendo dois primeiros valores do canal. A leitura causa um sincronismo. Ou seja, espera a goroutine executar
	fmt.Println("Logo apos a primeira leitura do channel")
	fmt.Println(a, b)

	fmt.Println(<-ch)
	fmt.Println("Logo apos a segunda leitura do channel")

	// <-ch se lermos de novo, teremos um erro pois não há mais nenhum valor para ler do canal.
}
