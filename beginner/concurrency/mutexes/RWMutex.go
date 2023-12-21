package main

import (
	"fmt"
	"sync"
)

/*
RW MUTEX
The standard library also exposes a sync.RWMutex

In addition to these methods:

Lock()
Unlock()
The sync.RWMutex also has these methods:

RLock()
RUnlock()
The sync.RWMutex can help with performance if we have a read-intensive process.
Many goroutines can safely read from the map at the same time (multiple Rlock() calls can happen simultaneously).
However, only one goroutine can hold a Lock() and all RLock()'s will also be excluded.

ASSIGNMENT
Let's update our same code from the last assignment, but this time we can speed it up by allowing readers to read from the map concurrently.

Update the val method to only lock the mutex for reading. Notice that if you run the code with a write lock it will block forever.

Lock(): only one go routine read/write at a time by acquiring the lock.

RLock(): multiple go routine can read(not write) at a time by acquiring the lock.

1) When a go-routine has already acquired a RLock(), can another go-routine acquire a Lock() for write or it has to wait until RUnlock() happens?

To acquire a Lock() for write it has to wait until RUnlock()
2) What happens when someone already acquired Lock() for map ,will other go-routine can still get RLock()

if someone X already acquired Lock(), then other go-routine to get RLock() will have to wait until X release lock (Unlock())
3) Assuming we are dealing with Maps here, is there any possibility of "concurrent read/write of Map" error can come?
*/

func main() {
	fmt.Println("main method")
	m := map[int]int{}
	mu := &sync.RWMutex{}

	go writeLoop(m, mu)

	go readLoop(m, mu)
	//go readLoop(m, mu)
	//go readLoop(m, mu)
	//go readLoop(m, mu)
	//go readLoop(m, mu)

	block := make(chan struct{})
	<-block
}

func writeLoop(m map[int]int, mu *sync.RWMutex) {
	for {
		fmt.Println("writing loop started...")
		for i := 0; i < 100; i++ {
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}
	}
}
func readLoop(m map[int]int, mu *sync.RWMutex) {
	fmt.Println("readloop starts")
	for {
		mu.RLock()
		for k, v := range m {
			fmt.Println(k, "", v)
		}
		defer mu.RUnlock()
	}

}
