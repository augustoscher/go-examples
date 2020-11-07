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
	// 1. struct to json
	p1 := Product{1, "Macbook", 1999.90, []string{"i5", "8gb ram"}}

	//retorna um slice de bytes e o erro
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// 2. json to struct
	var p2 Product
	jsonString := `{"id": 2,"name":"Macbook Pro","price":2999.9,"tags":["i7","16gb ram"]}`

	// Passa um slice de bytes com base na string json e a referência do objeto p2
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
	fmt.Println(p2.Tags[1])
}
