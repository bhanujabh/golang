package main

import "fmt"

// er at the end is used as convention
// give methods in interface 
// in go we don't have to explicitly mention where we are implementing the interface 
// if the method signatures present in an interface are also associated with some struct 
// then the compiler understands that the struct is using that interface 
// for ex. the pay method in razorpay 
type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter
}

// Open close principle
func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw.pay(amount)
	// stripePaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("making payment using stripe", amount)
// }

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway for testing purpose")
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string) {

}

func main() {
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}

	// fakeGw := fakepayment{}
	paypalGw := paypal{}
	newPayment := payment{
		gateway: paypalGw,
	}
	newPayment.makePayment(100)
}