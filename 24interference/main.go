// Interference in Golang
package main

import "fmt"

type paymenter interface {
	pay(amount float32) //We can also provide return value by pay(amount int) int
	refund(amount float32, price string)
}

type payment struct {
	gateway paymenter
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment using razor pay", amount)
}

func (r razorpay) refund(amount float32, price string) {
	fmt.Println("Payment done for this scenario")
}

func (p payment) makepayment(amount float32) {
	p.gateway.pay(amount)
	p.gateway.refund(amount, "done")
}

func main() {
	razorPayment := razorpay{}
	newPayment := payment{
		gateway: razorPayment,
	}
	newPayment.makepayment(100)
}
