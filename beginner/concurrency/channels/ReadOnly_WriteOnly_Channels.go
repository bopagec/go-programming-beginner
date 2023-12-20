package main

import (
	"fmt"
	"time"
)

/*
TICKERS
time.Tick() is a standard library function that returns a channel that sends a value on a given interval.
time.After() sends a value once after the duration has passed.
time.Sleep() blocks the current goroutine for the specified amount of time.

READ-ONLY CHANNELS
A channel can be marked as read-only by casting it from a chan to a <-chan type. For example:

func main(){
  ch := make(chan int)
  readCh(ch)
}

func readCh(ch <-chan int) {
  // ch can only be read from
  // in this function
}

WRITE-ONLY CHANNELS
The same goes for write-only channels, but the arrow's position moves.

func writeCh(ch chan<- int) {
  // ch can only be written to
  // in this function
}

ASSIGNMENT
Like all good back-end engineers, we frequently save backup snapshots of the Mailio database.

Complete the saveBackups function.

It should read values from the snapshotTicker and saveAfter channels simultaneously.

If a value is received from snapshotTicker, call takeSnapshot()
If a value is received from saveAfter, call saveSnapshot() and return from the function: you're done.
If neither channel has a value ready, call waitForData() and then time.Sleep() for 500 milliseconds. After all, we want to show in the logs that the snapshot service is running.
*/

func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot()
		case <-saveAfter:
			saveSnapshot()
			return // exit the infinite loop once saved.
		default:
			waitForData()
			time.Sleep(time.Millisecond * 500)
		}
	}

}

// don't touch below this line
func takeSnapshot() {
	fmt.Println("Taking a backup snapshot...")
}

func saveSnapshot() {
	fmt.Println("All backups saved!")
}

func waitForData() {
	fmt.Println("Nothing to do, waiting...")
}

func tester() {
	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	saveBackups(snapshotTicker, saveAfter)
	fmt.Println("===========================")
}

func main() {
	tester()
}
