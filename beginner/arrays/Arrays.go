package main

import "fmt"

/*
Arrays are fixed-size groups of variables of the same type.
The type [n]T is an array of n values of type T
To declare an array of 10 integers:

var myInts [10]int
*/
/*
ASSIGNMENT
When a message isn't responded to, we allow our clients to have up to 2 additional messages that are sent as nudging reminders.

getMessageWithRetries returns an array of 3 strings where index 0 is the first message.
If the first message isn't answered by the recipient, we send the second, if that one isn't answered then we send the third.

Update getMessageWithRetries to return the following 3 strings in an array.
*/

const (
	retry1 = "click here to sign up"
	retry2 = "pretty please click here"
	retry3 = "we beg you to sign up"
)

func main() {
	testSend("Bob", 0)
	testSend("Alice", 1)
	testSend("Mangalam", 2)
	testSend("Ozgur", 3)
}

func testSend(name string, doneAt int) {
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages := getMessageWithRetries()

	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf("sending: '%v'", msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded!")
		}
		if i == len(messages)-1 {
			fmt.Println("complete failure!")
		}
	}

}

func getMessageWithRetries() [3]string {
	return [3]string{retry1, retry2, retry3}
}
