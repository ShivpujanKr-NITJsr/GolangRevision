package main

import (
	"fmt"
	"time"
)

// orders structs

type customer struct {
	name  string
	email string
}

// composition
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nano second precision
	customer            // embedded struct
}

func newOrder(id string, amount float32, status string) *order {
	// initial setup goes here...
	myOrder := order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(), // set the createdAt field to the current time
	}
	return &myOrder // return the order struct
}

// receiver type
func (o *order) changeOrderStatus(newStatus string) {
	o.status = newStatus // change the status of the order
}

func (o order) getOrderStatus() string {
	return o.status // return the status of the order
}

func main() {

	// myOrder := order{
	// 	id:     "1",
	// 	amount: 100.50,
	// 	status: "Pending",
	// }

	// myOrder.createdAt = time.Now() // set the createdAt field to the current time
	// fmt.Println("Order struct:", myOrder) // prints the order struct

	// myOrder2 := order{
	// 	id:     "1",
	// 	amount: 100.50,
	// 	status: "Delivered",
	// 	createdAt: time.Now(), // set the createdAt field to the current time
	// }

	// myOrder.status = "Delivered" // change the status of the order
	// fmt.Println("Order struct 1:", myOrder) // prints the order struct 1
	// fmt.Println("Order struct 2:", myOrder2) // prints the order struct 2

	// myOrder := order{
	// 	id:     "1",
	// 	amount: 100.50,
	// 	status: "Pending",
	// }
	// myOrder.changeOrderStatus("Delivered")                                       // change the status of the order using receiver type method
	// fmt.Println("Order struct after change:", myOrder, myOrder.getOrderStatus()) // prints the order struct after change

	// myOrder := newOrder("1", 100.50, "Pending")
	// fmt.Println(myOrder.amount)

	// language := struct {
	// 	name	string
	// 	isGood	bool
	// }{"golang", true} // anonymous struct
	// fmt.Println("Language struct:", language) // prints the anonymous struct
	// newCustomer := customer{
	// 	name:  "John Doe",
	// 	email: "test@test.com",
	// }

	newOrder := order{
		id:     "1",
		amount: 100.50,
		status: "Pending",
		customer: customer{
			name:  "John Doe",
			email: "test@test.com",
		},
	}

	newOrder.customer.name = "Jane Doe" // change the name of the customer

	fmt.Println(newOrder)
}
