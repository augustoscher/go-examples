package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2 %s %s \n", p1, p2)
}

func f3() string {
	return "F3"
}

// Pure function: Receives a group of arguments, do somethin with params and maybe return value.
// Pure function never change any state outside of function scope/block
func f4(p1, p2 string, p3 int) string {
	return fmt.Sprintf("F4 %s - %s - %d", p1, p2, p3)
}

// Returning multiple values
func f5() (string, int64) {
	return "One", 2
}

// named return
func change(p1, p2 int) (xunda2 int, xunda1 int) {
	xunda2 = p2
	xunda1 = p1
	return // clean return
}

func main() {
	f1()
	f2("Oi,", "Tudo bem?")
	fmt.Println(f3())
	fmt.Println(f4("Oi,", "Tudo bem?", 1392234))
	message, value := f5()
	fmt.Printf("Message: %s - Value: %d\n", message, value)

	// valid:
	// _, r5 := f5() ignoring first return value
	// r5, _ := f5() ignoring second return value
	// f5() not expect for returns

	// not valid:
	// r5 := f5() not same returns are mapped

	xunda2, xunda1 := change(1000, 1001)
	fmt.Printf("Xunda2: %d - Xunda1: %d\n", xunda2, xunda1)
}
