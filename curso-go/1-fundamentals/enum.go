package main

import (
	"fmt"
)

type SubscriptionChangeEventEnum string

const (
	EventPaid SubscriptionChangeEventEnum = "PAID"
	EventOverdue SubscriptionChangeEventEnum = "OVERDUE"
	EventChangePlan SubscriptionChangeEventEnum = "CHANGE_PLAN"
)

type SubscriptionEvent struct {
	Event     SubscriptionChangeEventEnum `json:"event" binding:"required"`
	User      string `json:"user" binding:"required"`
	OrderCode string `json:"orderCode" binding:"required"`
}

func main() {

	es1 := SubscriptionEvent{EventPaid, "user", "orderCode"}
	es2 := SubscriptionEvent{EventOverdue, "user", "orderCode"}
	
	fmt.Printf("Object: %v\n", es1)
	fmt.Printf("Object: %v\n", es2)

	if (es1.Event == EventPaid) {
		fmt.Println("EventPaid on es1")
	}
}