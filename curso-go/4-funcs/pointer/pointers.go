package main

import "fmt"

func inc1(n int) {
	n++
}

//Ponteiro é representado por un *
func inc2(n *int) {
	// * usado para obter o valor ao qual o ponteiro aponta
	// *n retorna o valor e então incrementamos
	*n++

}

func main() {
	n := 1

	//passando por valor (valor de n não é alterado)
	inc1(n)
	fmt.Println(n) //1

	//passando o endereço de memória (ponteiro) para a inc2
	//passando por referência, onde o valor de n será alterado
	inc2(&n)
	fmt.Println(n) //2
}
