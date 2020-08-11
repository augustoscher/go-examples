package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha")

	fmt.Println(" Nova")
	fmt.Println("Linha")

	x := 3.14156
	xs := fmt.Sprint(x)
	// concatena string
	aux := "O valor de x é: " + xs
	fmt.Println(aux)
	fmt.Println("O valor de x é:", xs)
	fmt.Printf("O valor de x é: %f\n", x)
	fmt.Printf("O valor de x é: %.2f\n", x)

	a, b, c, d := 1, 1.9999, false, "opa"
	fmt.Printf("%d %f %t %s \n", a, b, c, d)
	fmt.Printf("%v %v %v %v \n", a, b, c, d)
}
