package main

import (
	"fmt"
	"math"
)

// interface is a collection of method signatures
// interfaces are not classes in Go,
// interfaces are not hierarchical by nature in Go,
// interfaces are much simpler in Go,
// interfaces define function signatures, but not underlying behaviour.
func main() {

	testShape(circle{
		radius: 10,
	})
	testShape(rectangle{
		width:  20,
		height: 10,
	})
}

func testShape(s shape) {
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

type shape interface {
	area() float64
	perimeter() float64
}

// because rectangle has implements both of methods in shape interface.
// we can think rectangle as a shape.
type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return r.height*2 + r.width*2
}

// because circle has implements both of methods in shape interface.
// we can think circle as a shape.
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
