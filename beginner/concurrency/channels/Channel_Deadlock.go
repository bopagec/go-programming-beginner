package main

// https://www.youtube.com/watch?v=VkGQFFl66X4
import "fmt"

func main() {
	dataChan := make(chan int)

	// sending data to the channel in a new thread
	// if we do not the sending data as a separate go routing it will create a deadlock
	go func() {
		for i := 0; i < 100; i++ {
			dataChan <- i
		}
		close(dataChan)
	}() // go routing is done by here

	// receiving data from the channel on main-thread
	// but, this runs forever waiting for data from the background thread on the main thread causing deadlock
	// for that reason, we must close the data channel once we finished sending the data in line no: 14
	for data := range dataChan {
		fmt.Printf("recieving data = %d\n", data)
	}
}
