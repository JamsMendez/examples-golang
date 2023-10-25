package main

import "fmt"

type PaymentMethod interface {
	ProcessPayment(amount float32)
}

type CreditCard struct {
}

func (cc *CreditCard) ProcessPayment(amount float32) {
	fmt.Printf("Processing credit card payment %.2f\n", amount)
}

type Paypal struct {
}

func (p *Paypal) ProcessPayment(amount float32) {
	fmt.Printf("Processing paypal payment %.2f\n", amount)
}

func main() {
	var paymentMethod PaymentMethod
	paymentMethod = &CreditCard{}
	paymentMethod.ProcessPayment(50.00)

	paymentMethod = &Paypal{}
	paymentMethod.ProcessPayment(15.00)
}
