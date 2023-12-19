package main

import "fmt"

/*
ASSIGNMENT
Complete the printReports function.

Call printCostReport once for each message.
Pass in an anonymous function as the costCalculator that returns an int equal to twice the length of the input message.
*/

func main() {
	test5([]string{
		"Here's Johnny!",
		"Go ahead, make my day",
		"You had me at hello",
		"There's no place like home",
	})

	test5([]string{
		"Hello, my name is Inigo Montoya. You killed my father. Prepare to die.",
		"May the Force be with you.",
		"Show me the money!",
		"Go ahead, make my day.",
	})
}

func test5(messages []string) {
	defer fmt.Println("====================================")
	printReports(messages)
}

func printReports(messages []string) {
	for _, msg := range messages {
		printCostReport(func(s string) int {
			return len(s) * 2
		}, msg)
	}
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
