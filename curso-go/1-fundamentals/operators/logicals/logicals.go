package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	buy50TV := trab1 && trab2
	buy32TV := trab1 != trab2 // ou exclusivo
	buyIceCream := trab1 || trab2

	return buy50TV, buy32TV, buyIceCream
}

func main() {
	tv50, tv32, iceCream := compras(true, true)
	fmt.Printf("TV50: %t, TV32: %t and IceCream: %t\n", tv50, tv32, iceCream)

	tv50, tv32, iceCream = compras(true, false)
	fmt.Printf("TV50: %t, TV32: %t and IceCream: %t\n", tv50, tv32, iceCream)

	tv50, tv32, iceCream = compras(false, true)
	fmt.Printf("TV50: %t, TV32: %t and IceCream: %t\n", tv50, tv32, iceCream)

	tv50, tv32, iceCream = compras(false, false)
	fmt.Printf("TV50: %t, TV32: %t and IceCream: %t\n", tv50, tv32, iceCream)
}
