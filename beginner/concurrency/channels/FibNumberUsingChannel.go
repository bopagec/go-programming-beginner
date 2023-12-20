package main

import (
	"fmt"
	"time"
)

/*
ASSIGNMENT
It's that time again, Mailio is hiring and we've been assigned to do the interview.
For some reason, the Fibonacci sequence is Mailio's interview problem of choice.
We've been tasked with building a small toy program we can use in the interview.

Complete the concurrentFib function. It should:

Create a new channel of ints
Call fibonacci in a goroutine, passing it the channel and the number of Fibonacci numbers to generate, n
Use a range loop to read from the channel and print out the numbers one by one, each on a new line
*/

func concurrentFib(n int) {
	channel := make(chan int)
	go func() {
		// this will run on a different go-routine
		fibonacci(n, channel)
		// when fibonacci is done it will close the channel.
	}()
	// reader goes here, not inside the go func.
	// this will run on main go-routine
	for val := range channel {
		fmt.Println(val)
	}
}

// don't touch below this line

func fibTest(n int) {
	fmt.Printf("Printing %v numbers...\n", n)
	concurrentFib(n)
	fmt.Println("==============================")
}

func main() {
	fibTest(10)
	fibTest(5)
	fibTest(20)
	fibTest(13)
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}
