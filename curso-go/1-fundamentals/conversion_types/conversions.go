package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int para float
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	//float para int
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println("Nota:", notaFinal) // 6

	// string
	fmt.Println("Teste " + string(123)) //string(123) retorna o valor ascii

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123)) //Teste 123

	// string para int
	num, _ := strconv.Atoi("10")
	fmt.Println("str to int", num)

	// string para boolean
	b, _ := strconv.ParseBool("true") // funciona com "1" tb
	if b {
		fmt.Println("str to bool", b)
	}
}
