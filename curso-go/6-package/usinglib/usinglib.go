package main

import (
	"fmt"

	"github.com/augustoscher/goarea/test"

	"github.com/augustoscher/goarea"
)

func main() {
	// Installing package:
	// go get -u github.com/augustoscher/goarea

	// Just using area.CircunferenceArea go imports package created on gopath folder
	// Is not necessary compile the package file.
	// When we run it at the first time, he's automatically compiled
	fmt.Println(goarea.CircunferenceArea(6.0))

	fmt.Println(goarea.RectangleArea(3.0, 5.4))

	test.PrintValue("Xunda")

	// Not visible because init with _
	// fmt.Println(goarea._EquilateralTriangleArea(3.0, 4.0))
}
