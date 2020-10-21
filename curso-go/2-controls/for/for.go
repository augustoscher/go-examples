package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("while way")
	i := 1
	for i <= 10 {
		fmt.Printf("i: %d\n", i)
		i++
	}

	fmt.Println("")
	fmt.Println("Tradicional")
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println("")

	fmt.Println("")
	fmt.Println("Mixing")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Numero %d é par\n", i)
		} else {
			fmt.Printf("Numero %d é impar\n", i)
		}
	}

	fmt.Println("")
	fmt.Println("Loop infinito")
	for {
		fmt.Println("Oi...")
		time.Sleep(time.Second)
	}

}
