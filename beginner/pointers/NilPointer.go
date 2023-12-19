package main

import (
	"fmt"
	"strings"
)

/*
ASSIGNMENT
Complete the removeProfanity function.

It should use the strings.ReplaceAll function to replace all instances of the following words in the input message with asterisks.

"dang" -> "****"
"shoot" -> "*****"
"heck" -> "****"
It should mutate the value in the pointer and return nothing.

Do not alter the function signature.
*/

// we want to mutate the *string without returning the changed string as return value
// see, there is no return string in the function
func removeProfanity_2(message *string) {
	if message == nil {
		return
	}
	messageVal := *message // de-reference the message pointer to value

	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")

	*message = messageVal // assign back the new value to the de-referenced message value
}

// don't touch below this line
// if we don't handle nil pointers it will panic as below
// panic: runtime error: invalid memory address or nil pointer dereference
func test2(messages []string) {
	for _, message := range messages {
		if message == "" {
			removeProfanity_2(nil)
			fmt.Println("nil message detected")
		} else {
			removeProfanity_2(&message)
			fmt.Println(message)
		}
	}
}

func main() {
	messages1 := []string{
		"well shoot, this is awful",
		"",
		"dang robots",
		"dang them to heck",
		"",
	}

	messages2 := []string{
		"well shoot",
		"",
		"Allan is going straight to heck",
		"dang... that's a tough break",
		"",
	}

	test2(messages1)
	test2(messages2)
}
