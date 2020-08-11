package main

import "fmt"

func main() {
	fmt.Printf(getName("Augusto"))
	testing()
	loopTest()
}

func getName(name string) string {
	return "Hello, " + name + "!"
}

func testing() {
	fmt.Print("\n")
	slice := []int{12, 1231, 12312}
	fmt.Print(slice)

	fmt.Print("\n")
	slice2 := []byte("Hello")
	fmt.Print(slice2)
}

func loopTest() {
	fmt.Print("\n")
	entry := []string{"Jack", "John", "Jones"}
	for i, val := range entry {
		fmt.Printf("At position %d, the character %s is present\n", i, val)
	}
}
