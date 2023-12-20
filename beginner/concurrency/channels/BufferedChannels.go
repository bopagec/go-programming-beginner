package main

import "fmt"

/*
BUFFERED CHANNELS
Channels can optionally be buffered.

CREATING A CHANNEL WITH A BUFFER
You can provide a buffer length as the second argument to make() to create a buffered channel:

ch := make(chan int, 100)
Copy icon
Sending on a buffered channel only blocks when the buffer is full.

Receiving blocks only when the buffer is empty.
ASSIGNMENT
We want to be able to send emails in batches.
A writing goroutine will write an entire batch of email messages to a buffered channel,
and later, once the channel is full, a reading goroutine will read all of the messages from the channel and send them out to our clients.

Complete the addEmailsToQueue function.
It should create a buffered channel with a buffer large enough to store all of the emails it's given.
It should then write the emails to the channel in order, and finally return the channel.
*/

func addEmailsToQueue(emails []string) chan string {
	// without having the buffer size, the go routine will get blocked forever.
	// because there is no reader on the other end to consume the channel.
	// there is a ready on the other end which runs on the same go routine
	emailsToSend := make(chan string, len(emails))
	for _, item := range emails {
		emailsToSend <- item
	}
	return emailsToSend
}
func sendEmails(bathSize int, channel chan string) {
	for i := 0; i < bathSize; i++ {
		item := <-channel
		fmt.Printf("sending email: %s\n ", item)
	}
}

func testBufferedChannel(emails ...string) {
	fmt.Printf("Adding %v emails to queue...\n", len(emails))
	ch := addEmailsToQueue(emails)
	fmt.Println("Sending emails...")
	sendEmails(len(emails), ch)
	fmt.Println("==========================================")
}

func main() {
	testBufferedChannel("Hello John, tell Kathy I said hi", "Whazzup bruther")
	testBufferedChannel("I find that hard to believe.", "When? I don't know if I can", "What time are you thinking?")
	testBufferedChannel("She says hi!", "Yeah its tomorrow. So we're good.", "Cool see you then!", "Bye!")
}
