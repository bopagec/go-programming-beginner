package main

/*
Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array.
If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller,
analogous to passing a pointer to the underlying array.

A Read function can therefore accept a slice argument rather than a pointer and a count;
the length within the slice sets an upper limit of how much data to read
*/
import (
	"errors"
	"fmt"
)

/*
99 times out of 100 you will use a slice instead of an array when working with ordered lists.
Arrays are fixed in size. Once you make an array like [10]int you can't add an 11th element.
A slice is a dynamically-sized, flexible view of the elements of an array.
Slices always have an underlying array, though it isn't always specified explicitly.
To explicitly create a slice on top of an array we can do:

primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4] // inclusive, exclusive
// mySlice = {3, 5, 7}

ASSIGNMENT
Retries are a premium feature now! Textio's free users only get 1 retry message, while pro members get an unlimited amount.

Complete the getMessageWithRetriesForPlan function. It takes a plan variable as input, and you've been provided with constants for the plan types at the top of the program.

If the plan is a pro plan, return all the strings from getMessageWithRetries().
If the plan is a free plan, return the first 2 strings from getMessageWithRetries().
If the plan isn't either of those, return an error that says unsupported plan.
*/
func main() {
	test2("Ozgur", 3, planFree)
	test2("Jeff", 3, planPro)
	test2("Sally", 2, planPro)
	test2("Sally", 3, "no plan")
}
func test2(name string, doneAt int, plan string) {
	defer fmt.Println("=====================================")
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages, err := getMessageWithRetriesForPlan(plan)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`sending: "%v"`, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("no response")
		}
	}
}

const (
	planFree = "free"
	planPro  = "pro"
)
const (
	retry_1 = "click here to sign up"
	retry_2 = "pretty please click here"
	retry_3 = "we beg you to sign up"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessage_With_Retries()

	if plan == planPro {
		return allMessages[0:3], nil
	}
	if plan == planFree {
		return allMessages[0:2], nil
	}
	return nil, errors.New("unsupported plan")
}
func getMessage_With_Retries() [3]string {
	return [3]string{retry_1, retry_2, retry_3}
}
