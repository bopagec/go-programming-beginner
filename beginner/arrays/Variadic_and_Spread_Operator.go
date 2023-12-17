package main

import "fmt"

/*
VARIADIC
Many functions, especially those in the standard library, can take an arbitrary number of final arguments. This is accomplished by using the "..." syntax in the function signature.
A variadic function receives the variadic arguments as a slice.

SPREAD OPERATOR
The spread operator allows us to pass a slice into a variadic function. The spread operator consists of three dots following the slice in the function call.
*/

/*
ASSIGNMENT
We need to sum up the costs of all individual messages so we can send an end-of-month bill to our customers.

Complete the sum function so that returns the sum of all of its inputs.

NOTE
Make note of how the variadic inputs and the spread operator are used in the test suite.
*/
func main() {
	test3(1.0, 2.0, 3.0)
	test3(1.0, 2.0, 3.0, 4.0, 5.0)
	test3(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
	test3(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0)
}

func test3(nums ...float64) { // variadic function
	total := sum(nums...) // spread operator
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %.2f\n", total)
	fmt.Println("===== END REPORT =====")
}

func sum(nums ...float64) float64 { // variadic function
	total := 0.0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}
