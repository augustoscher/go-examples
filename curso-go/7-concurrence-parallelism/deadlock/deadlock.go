package main

import (
	"fmt"
	"time"
)

func doSomething(c chan int) {
	time.Sleep(time.Second)
	c <- 1 //operação bloqueante
	fmt.Println("Só será printado depois que for lido...")
}

func main() {
	c := make(chan int) //canal sem buffer

	go doSomething(c)

	fmt.Println(<-c) //operacao bloqueante
	fmt.Println("Lido")
	fmt.Println(<-c) //deadlock
	fmt.Println("Não imprime pois ocorrerá erro na linha acima")
}
