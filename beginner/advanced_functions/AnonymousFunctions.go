package main

import "fmt"

/*
ANONYMOUS FUNCTIONS
Anonymous functions are true to form in that they have no name.
We've been using them throughout this chapter, but we haven't really talked about them yet.

Anonymous functions are useful when defining a function that will only be used once or to create a quick closure.
*/

// doMath accepts a function that converts one int into another
// and a slice of ints. It returns a slice of ints that have been
// converted by the passed in function.
func doMath(f func(int) int, nums []int) []int {
	var results []int
	for _, n := range nums {
		results = append(results, f(n))
	}
	return results
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// Here we define an anonymous function that doubles an int
	// and pass it to doMath
	allNumsDoubled := doMath(func(x int) int {
		return x + x
	}, nums)

	fmt.Println(allNumsDoubled)
	// prints:
	// [2 4 6 8 10]
}
