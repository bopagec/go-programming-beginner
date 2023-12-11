package main

import "fmt"

// in Go, there are no classes, instead it is called Structs where we structure the data.
// Structs ( Or rather Structures)
// we use Structs in Go to represent structured data,
// for example if you want to represent a car we can create a struct with it's key value pair types
// In Go, you will often use a struct to represent information that you would have used a dictionary for in Python, or an object literal in JavaScript
func main() {
	// instantiate a empty message with default value
	msg := messageToSend{}

	test(messageToSend{
		phoneNumber: "07983397030",
		message:     "Hello World",
	})
	test(msg)

	sender := user{name: "Charatha", number: "07983397030"}
	receiver := user{name: "Nush", number: "07882421145"}
	userMessage := userMessage{
		message:  "Hello World",
		sender:   sender,
		receiver: receiver,
	}

	testUserMessage(userMessage)

}

func testUserMessage(message userMessage) {
	fmt.Printf("Sending user message: '%s' from : '%s' (%s) to : '%s' (%s)", message.message, message.sender.name, message.sender.number, message.receiver.name, message.receiver.number)
}

func test(message messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", message.message, message.phoneNumber)
	fmt.Println("===========================================")
}

type messageToSend struct {
	message     string
	phoneNumber string
}

type userMessage struct {
	message  string
	sender   user
	receiver user
}
type user struct {
	name   string
	number string
}
