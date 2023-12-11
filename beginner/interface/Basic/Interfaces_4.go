package main

import "fmt"

func main() {
	e := email{
		isSubscribed: false,
		body:         "Hello World!",
	}

	emailTest(e, e)
	e = email{
		isSubscribed: true,
		body:         "Hello World!",
	}

	emailTest(e, e)
}

func emailTest(e expenses, p printer) {
	fmt.Printf("Printing with cost: $%.2f\n", e.cost())
	p.print()
	fmt.Println("==============================")
}

type expenses interface {
	cost() float64
}

type printer interface {
	print()
}
type email struct {
	isSubscribed bool
	body         string
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05 * float64(len(e.body)) // in Go, we can't multiply int with float unless we cast to float64
	}
	return 0.01 * float64(len(e.body))
}

func (e email) print() {
	fmt.Println(e.body)
}
