package main

import "fmt"

func main() {
	// if statements in Go do not use parentheses around the condition:
	messageLen := 10
	maxMessageLen := 20

	fmt.Println("trying to send message of length: ", messageLen)

	if messageLen > 1 && messageLen < maxMessageLen {
		fmt.Println("Message sent")
	} else if messageLen < 1 {
		fmt.Println("Message length too small")
	} else {
		fmt.Println("Message not sent")
	}

	// another way of writing if condition by removing the length variable scope
	// the length variable scope is only accessible by only the if condition
	// we can make sure length variable is never being used by any other code for safety.

	if length := getLength("charatha@gmail.com"); length > 1 {
		fmt.Println("email length is, ", length)
	}

}

func getLength(email string) int {
	return len(email)
}
