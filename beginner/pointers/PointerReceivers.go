package main

import "fmt"

/*
POINTER RECEIVERS
A receiver type on a method can be a pointer.

Methods with pointer receivers can modify the value to which the receiver points.
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
*/

type car struct {
	color string
}

type bike struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}
func (b bike) setColor(color string) {
	b.color = color
}
func main() {
	c := car{color: "white"}
	b := bike{color: "white"}

	c.setColor("blue")
	b.setColor("blue")

	fmt.Println(c.color) // blue
	fmt.Println(b.color) // white
}
