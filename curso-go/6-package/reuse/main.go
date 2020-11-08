package main

import (
	"fmt"

	"github.com/augustoscher/area"
)

func main() {
	// Just using area.CircunferenceArea go imports package created on gopath folder
	// Is not necessary compile the package file.
	// When we run it at the first time, he's automatically compiled
	fmt.Println(area.CircunferenceArea(6.0))

	fmt.Println(area.RectangleArea(3.0, 5.4))

	// Not visible
	// fmt.Println(area._EquilateralTriangleArea(3.0, 4.0))
}
