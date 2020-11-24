package main

import "fmt"

func main() {
	//Channel é um tipo da linguagem.

	ch := make(chan int, 1)
	ch <- 1 //escreve um valor inteiro em um canal
	<-ch    //ler dados de um canal

	ch <- 3
	fmt.Println(<-ch)
}
