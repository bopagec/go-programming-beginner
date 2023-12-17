package main

import "fmt"

/*
MAKE
Most of the time we don't need to think about the underlying array of a slice. We can create a new slice using the make function:

Cap(slice) // this will return the capacity of the slice

//func make([]T, len, cap) []T
mySlice := make([]int, 5, 10)

//the capacity argument is usually omitted and defaults to the length
mySlice := make([]int, 5)

slice literals, note that the array brackets do not have 3 in them.
mySlice := []string{"one", "two", "three"}

ASSIGNMENT
We send a lot of text messages at Textio, and our API is getting slow and unresponsive.

If we know the rough size of a slice before we fill it up,
we can make our program faster by creating the slice with that size ahead of time
so that the Go runtime doesn't need to continuously allocate new underlying arrays of larger and larger sizes. By setting the length,
the slice can still be resized later, but it means we can avoid all the expensive resizing since we know that we'll need.

Complete the getMessageCosts() function. It takes a slice of messages and returns a slice of message costs.

Preallocate a slice for the message costs of the same length as the messages slice.
Fill the costs slice with costs for each message.
The cost in the cost slice should correspond to the message in the messages slice at the same index.
The cost of a message is the length of the message multiplied by 0.01.
*/
func main() {
	test([]string{
		"Welcome to the movies!",
		"Enjoy your popcorn!",
		"Please don't talk during the movie!",
	})
	test([]string{
		"I don't want to be here anymore",
		"Can we go home?",
		"I'm hungry",
		"I'm bored",
	})
	test([]string{
		"Hello",
		"Hi",
		"Hey",
		"Hi there",
		"Hey there",
		"Hi there",
		"Hello there",
		"Hey there",
		"Hello there",
		"General Kenobi",
	})
}

func test(messages []string) {
	costs := getMessageCosts(messages)
	fmt.Println("Messages:")
	for i := 0; i < len(messages); i++ {
		fmt.Printf(" - %v\n", messages[i])
	}
	fmt.Println("Costs:")
	for i := 0; i < len(costs); i++ {
		fmt.Printf(" - %.2f\n", costs[i])
	}
	fmt.Println("===== END REPORT =====")
}

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		cost := float64(len(messages[i])) * 0.01
		costs[i] = cost
	}
	return costs
}
