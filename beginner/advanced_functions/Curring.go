package main

import "fmt"

/*
CURRYING
Function currying is the practice of writing a function that takes a function (or functions) as input, and returns a new function.
*/
func main() {
	squareFunc := selfMath(multiply)
	doubleFunc := selfMath(addition)

	fmt.Println(squareFunc(5))

	fmt.Println(doubleFunc(5))
}

func multiply(x, y int) int {
	return x * y
}

func addition(x, y int) int {
	return x + y
}

func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}
