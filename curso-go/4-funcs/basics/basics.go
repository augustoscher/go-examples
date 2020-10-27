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

func main() {
	f1()
	f2("Oi,", "Tudo bem?")
	fmt.Println(f3())
	fmt.Println(f4("Oi,", "Tudo bem?", 1392234))
}
