package main

/*
CHANNELS
Empty structs are often used as tokens in Go programs. In this context, a token is a unary value. In other words, we don't care what is passed through the channel. We care when and if it is passed.

We can block and wait until something is sent on a channel using the following syntax

<-ch

This will block until it pops a single item off the channel, then continue, discarding the item.

ASSIGNMENT
Our Mailio server isn't able to boot up until it receives the signal that its databases are all online, and it learns about them being online by waiting for tokens (empty structs) on a channel.

Complete the waitForDbs function. It should block until it receives numDBs tokens on the dbChan channel. Each time it reads a token the getDatabasesChannel goroutine will print a message to the console for you.
*/
import "fmt"

// Channel we use to send and receive data
// Channel must have both portal open,
// if only 1 portal open, it will cause deadlock
func main() {
	// making an empty channel
	dataChan := make(chan int)

	// send data to the channel from the main-thread
	// main thread will block dataChan forever causing the deadlock

	// with go routing - if we send data to the channel from a different go routine it will work
	go func() {
		dataChan <- 5
	}()

	// receive data from channel from the main-thread
	// with go routing when receiving data - the main thread can continue to receive data, so it will work
	val := <-dataChan

	// print value from the main-thread
	fmt.Println(val)

}
