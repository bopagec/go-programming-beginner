package main

import "fmt"

// loops in Go is much similar to loops in other languages
// only difference is it does not have parenthesis

/*
ASSIGNMENT
At Textio we have a dynamic formula for determining how much a batch of bulk messages costs to send.

COMPLETE THE BULKSEND() FUNCTION
This function should return the total cost (as a float64) to send a batch of numMessages messages. Each message costs 1.0, plus an additional fee. The fee structure is:

1st message: 1.0 + 0.00
2nd message: 1.0 + 0.01
3rd message: 1.0 + 0.02
4th message: 1.0 + 0.03
...
Use a loop to calculate the total cost and return it.
*/
func main() {
	test1(10)
	test1(20)
	test1(30)
	test1(40)
	test1(50)
}

func test1(numMessages int) {
	fmt.Printf("Sending %v messages \n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("===============================================================")
}

func bulkSend(messages int) float64 {
	cost := 0.0
	for i := 0; i < messages; i++ {
		cost += 1 + (0.01 * 1)
	}
	return cost
}
