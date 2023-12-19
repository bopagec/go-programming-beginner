package main

import "fmt"

func main() {
	var x int = 50
	var y *int = &x
	*y = 100

	// what is the value of *y after the code executes
	fmt.Println(*y) // 100

	// what is the value of y  after the code executes
	fmt.Println(y) // memory addr of the x

	// what is the value of x  after the code executes
	fmt.Println(x) // 100, because *y changed the value to 100
}
