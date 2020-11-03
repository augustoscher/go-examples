package main

import "fmt"

type OrderItem struct {
	productId int
	qtd       float64
	price     float64
}

type Order struct {
	userId int
	items  []OrderItem
}

func (o Order) total() float64 {
	total := 0.00
	for _, item := range o.items {
		total += item.qtd * item.price
	}
	return total
}

func main() {
	order := Order{
		userId: 1,
		items: []OrderItem{
			OrderItem{
				productId: 1,
				qtd:       2.00,
				price:     1499.99,
			},
			OrderItem{
				productId: 2,
				qtd:       2.00,
				price:     900.00,
			},
		},
	}

	fmt.Println("Order: ", order)
	fmt.Println("Total: ", order.total())
}
