package main

import "fmt"


type paymenter interface {
	pay(amount float32)
	refund(amount float32,account string)
}


type payment struct {
	// gateway stripe
	// gateway razorpay
	gateway paymenter 
}

// open close principle
func (p payment) makePayment(amount float32) {
	// process payment logic here
	// razorpayPayment := razorpay{}
	// razorpayPayment.pay(amount)

	// stripePayment := stripe{}
	// stripePayment.pay(amount)
	p.gateway.pay(amount)

}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// process payment logic for Razorpay here
	fmt.Println("Payment made using Razorpay for amount:", amount)
}

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	// process payment logic for Stripe here
// 	fmt.Println("Payment made using Stripe for amount:", amount)
// }

type fakePayment struct{}

func (f fakePayment) pay(amount float32) {
	// process payment logic for fake payment here
	fmt.Println("Fake payment made for amount:", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	// process payment logic for PayPal here
	fmt.Println("Payment made using PayPal for amount:", amount)
}

func (p paypal) refund(amount float32, account string) {
	// process refund logic for PayPal here
	fmt.Println("Refund processed using PayPal for amount:", amount, "to account:", account)
}

func main() {
	// stripePayment := stripe{}
	// razorpayPayment := razorpay{}

	// fakeGW := fakePayment{}

	paypalGW := paypal{}

	newPayment := payment{
		// gateway: stripePayment, 
		// gateway: razorpayPayment, 
		// gateway: fakeGW, 
		gateway: paypalGW, 
	}
	newPayment.makePayment(100.50)

}
