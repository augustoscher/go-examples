package main

import (
	"fmt"
	"time"
)

// A escrita de dados é bufferizada.
// Operação não bloqueante.
// Quando temos um buffer de 10 posições, só vai bloquear depois que as 10 tiverem ocupadas.
// Quando mandar a 11º informação, vai ficar esperando alguem ler pelo menos 1 info para escrever a proxima.

func doSomething(ch chan int) {
	fmt.Println("doing something")
	ch <- 1
	ch <- 2
	ch <- 3
	//aqui o buffer foi completado
	ch <- 4 // vai esperar um valor ser consumido (linha 30)
	fmt.Println("executará somente se um valor do channel for lido")
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3) //criando canal com buffer de 3 posições
	go doSomething(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)

	time.Sleep(time.Second * 2) //só para esperar o buffer processar
	fmt.Println("fim")
}
