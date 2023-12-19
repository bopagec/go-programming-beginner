package main

import "fmt"

/*
POINTER RECEIVER CODE
Methods with pointer receivers don't require that a pointer is used to call the method.
The pointer will automatically be derived from the value.

type circle struct {
*/
type circle struct {
	x      int
	y      int
	radius int
}

func (c *circle) grow() {
	c.radius *= 2
}

func main() {
	// we do not have to use c: &circle to create this circle as a pointer
	c := circle{
		x:      1,
		y:      2,
		radius: 4,
	}

	// notice c is not a pointer in the calling function
	// but the method still gains access to a pointer to c
	c.grow()
	fmt.Println(c.radius)
	// prints 8
}
