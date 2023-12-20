package main

/*
CHANNELS REVIEW
Here are a few extra things you should understand about channels from Dave Cheney's awesome article.

1) A SEND TO A NIL CHANNEL BLOCKS FOREVER
var c chan string // c is nil
c <- "let's get started" // blocks

2) A RECEIVE FROM A NIL CHANNEL BLOCKS FOREVER
var c chan string // c is nil
fmt.Println(<-c) // blocks

3) A SEND TO A CLOSED CHANNEL PANICS
var c = make(chan int, 100)
close(c)
c <- 1 // panic: send on closed channel

4) A RECEIVE FROM A CLOSED CHANNEL RETURNS THE ZERO VALUE IMMEDIATELY
var c = make(chan int, 100)
close(c)
fmt.Println(<-c) // 0
*/
