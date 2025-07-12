package main

import "fmt"

// enumerated types  ,no inbuilt but using const-custom type

// iota -auto incremented value
// iota starts from 0 and increments by 1 for each constant in the block

// type OrderStatus int
type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to:", status)
}

func main() {
	changeOrderStatus(Received)
}
