package main

import "fmt"

/*
Most programming languages have a concept of a while loop.
Because Go allows for the omission of sections of a for loop, a while loop is just a for loop that only has a CONDITION.

	for CONDITION {
	  // do some stuff while CONDITION is true
	}
*/ /*
ASSIGNMENT
We have an interesting new cost structure from our SMS vendor.
They charge exponentially more money for each consecutive text we send!
Let's write a function that can calculate how many messages we can send in a given batch given a costMultiplier and a maxCostInPennies.

In a nutshell, the first message costs a penny, and each message after that costs the same as the previous message multiplied by the costMultiplier. That gets expensive!

There is an infinite loop in the code! Add a condition to the for loop to fix the bug.
The loop should break as soon as the actualCostInPennies is greater than the maxCostInPennies.
*/

func main() {
	test3(1.1, 5)
	test3(1.3, 10)
	test3(1.35, 25)
}

func test3(costMultiplier float64, maxCostInPennies int) {
	maxMessagesToSend := getMaxMessagesToSend(costMultiplier, maxCostInPennies)
	fmt.Printf("Multiplier: %v\n",
		costMultiplier,
	)
	fmt.Printf("Max cost: %v\n",
		maxCostInPennies,
	)
	fmt.Printf("Max messages you can send: %v\n",
		maxMessagesToSend,
	)
	fmt.Println("====================================")
}

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	for actualCostInPennies <= float64(maxCostInPennies) {
		actualCostInPennies *= costMultiplier
		maxMessagesToSend++
	}
	if actualCostInPennies > float64(maxCostInPennies) {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}
