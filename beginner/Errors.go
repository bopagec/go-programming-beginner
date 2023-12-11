package main

import (
	"fmt"
)

// Error in go are just values
// an Error is any type that implements the simple built-in error interface
// the error interface has a single method called Error() string method to describe what went wrong.
// if error is nil that means everything went fine and nothing went wrong.

// defer keyword
/*
The defer keyword in Golang is a powerful tool that can simplify
error handling and ensure that certain operations are always performed before a function returns.
By using the defer keyword, we can ensure that cleanup functions are executed,
even in the event of a panic.
This can help make code more robust and improve the overall reliability of our programs.
*/
func main() {
	testSendSMS(
		"Thanks for coming in to our flower shop today!",
		"We hope you enjoyed your gift.",
	)
	testSendSMS(
		"Thanks for joining us!",
		"Have a good day.",
	)
	testSendSMS(
		"Thank you.",
		"Enjoy!",
	)
	testSendSMS(
		"We loved having you in!",
		"We hope the rest of your evening is absolutely fantastic.",
	)
}

/*
ASSIGNMENT
We offer a product that allows businesses that use Textio to send pairs of messages to couples. It is mostly used by flower shops and movie theaters.
Complete the sendSMSToCouple function. It should send 2 messages, first to the customer, then to the customer's spouse.
Use sendSMS() to send the msgToCustomer. If an error is encountered, return 0.0 and the error.
Do the same for the msgToSpouse
If both messages are sent successfully, return the total cost of the messages added together.
When you return a non-nil error in Go, it's conventional to return the "zero" values of all other return values.
*/
func sendMsgToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	custCost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	spouseCost, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}

	return custCost + spouseCost, nil

}

func testSendSMS(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("================")
	fmt.Println("Message for Customer, ", msgToCustomer)
	fmt.Println("Message for Spouse, ", msgToSpouse)
	totalCost, err := sendMsgToCouple(msgToCustomer, msgToSpouse)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("Total cost: $%.4f\n", totalCost)
}

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const constPerChar = .0002

	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("Can't send texts over %v characters", maxTextLen)
	}

	return constPerChar * float64(len(message)), nil
}
