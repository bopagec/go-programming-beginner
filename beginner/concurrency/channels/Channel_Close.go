package main

import (
	"fmt"
	"time"
)

/*
CLOSING CHANNELS IN GO
Channels can be explicitly closed by a sender:

ch := make(chan int)

// do some stuff with the channel

close(ch)

CHECKING IF A CHANNEL IS CLOSED
Similar to the ok value when accessing data in a map,
receivers can check the ok value when receiving from a channel to test if a channel was closed.

v, ok := <-ch

ok is false if the channel is empty and closed.

DON'T SEND ON A CLOSED CHANNEL
Sending on a closed channel will cause a panic.
A panic on the main goroutine will cause the entire program to crash, and a panic in any other goroutine will cause that goroutine to crash.

Closing isn't necessary.
There's nothing wrong with leaving channels open, they'll still be garbage collected if they're unused.
You should close channels to indicate explicitly to a receiver that nothing else is going to come across.

ASSIGNMENT
At Mailio we're all about keeping track of what our systems are up to with great logging and telemetry.

The sendReports function sends out a batch of reports to our clients and reports back how many were sent across a channel.
It closes the channel when it's done.

Complete the countReports function. It should:

Use an infinite for loop to read from the channel:
If the channel is closed, break out of the loop
Otherwise, keep a running total of the number of reports sent
Return the total number of reports sent
*/
func countReports(numSentCh chan int) int {
	total := 0

	/*	we can avoid having the infinite loop by using the range keyword to loop the channel
		it will automatically stop looping when the channel is closed.
		but, the assignment was about checking if the channel is closed or not.
		for count := range numSentCh {
			total += count
		}
	*/
	for true {
		count, ok := <-numSentCh // make sure the channel is open or closed. if closed break the infinite loop
		if !ok {
			break
		}
		total += count
	}
	return total
}

func testSendReport(batchSize int) {
	channel := make(chan int)
	go sendReports(batchSize, channel)

	fmt.Println("Start counting...")
	numReports := countReports(channel)
	fmt.Printf("%v reports sent!\n", numReports)
	fmt.Println("=======================================")
}
func main() {
	testSendReport(3)
	testSendReport(4)
	testSendReport(5)
	testSendReport(6)
}

func sendReports(batchSize int, channel chan int) {
	for i := 0; i < batchSize; i++ {
		numReports := i*23 + 32%17
		channel <- numReports
		fmt.Printf("Sent batch of %v reports\n", numReports)
		time.Sleep(time.Millisecond * 10)
	}
	close(channel)

}
