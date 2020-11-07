package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to json
	p1 := Product{1, "Macbook", 1999.90, []string{"i5", "8gb ram"}}

	//retorna um slice de bytes e o erro
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))
}
