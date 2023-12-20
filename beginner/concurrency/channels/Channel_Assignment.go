package main

import (
	"fmt"
	"time"
)

/*
ASSIGNMENT
Our Mailio server isn't able to boot up until it receives the signal that its databases are all online, and it learns about them being online by waiting for tokens (empty structs) on a channel.

Complete the waitForDbs function.
It should block until it receives numDBs tokens on the dbChan channel.
Each time it reads a token the getDatabasesChannel goroutine will print a message to the console for you.
*/

func waitForDbs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		token := <-dbChan
		fmt.Printf("received token: %v\n", token)
	}
}

func testDBs(numDBs int) {
	dbChan := getDatabasesChannel(numDBs)
	fmt.Printf("Waiting for %v databases...\n", numDBs)
	waitForDbs(numDBs, dbChan)
	time.Sleep(time.Second)
	fmt.Println("All databases are online!")
	fmt.Println("===========================")
}
func main() {
	testDBs(3)
	testDBs(4)
	testDBs(5)
}

func getDatabasesChannel(numDBs int) chan struct{} {
	chanel := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			chanel <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}()

	return chanel
}
