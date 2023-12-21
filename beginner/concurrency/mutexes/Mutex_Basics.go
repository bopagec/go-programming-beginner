package main

/*
MUTEXES IN GO
Mutexes allow us to lock access to data. This ensures that we can control which goroutines can access certain data at which time.

Go's standard library provides a built-in implementation of a mutex with the sync.Mutex type and its two methods:

.Lock()
.Unlock()

We can protect a block of code by surrounding it with a call to Lock and Unlock as shown on the protected() method below.

It's good practice to structure the protected code within a function so that defer can be used to ensure that we never forget to unlock the mutex.

func protected(){
    mu.Lock()
    defer mu.Unlock()
    // the rest of the function is protected
    // any other calls to `mu.Lock()` will block
}

Mutexes are powerful. Like most powerful things, they can also cause many bugs if used carelessly.

MAPS ARE NOT THREAD-SAFE
Maps are not safe for concurrent use! If you have multiple goroutines accessing the same map, and at least one of them is writing to the map, you must lock your maps with a mutex.
*/
