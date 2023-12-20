package main

// https://www.youtube.com/watch?v=VkGQFFl66X4 ( 22 min )
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func simulateWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main() {
	// creating the channel
	dataChan := make(chan int)

	//sending data to th channel
	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := simulateWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	// receiving data from the channel in the main thread, one single portal to receive data from different go routines
	for data := range dataChan {
		fmt.Printf("data received: %d\n", data)
	}
}
