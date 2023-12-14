package main

import (
	"errors"
	"fmt"
)

/*
The Go standard library provides an "errors" package that makes it easy to deal with error.

# Errors.New() function can simply create an error with a custom error message as below

var err error = errors.New("something went wrong!")
*/
func main() {
	testDivider(10, 0)
	testDivider(10, 2)
	testDivider(15, 30)
	testDivider(6, 3)
}

type divideError struct {
	dividend float64
}

func (e divideError) Error() string {
	return fmt.Sprintf("%.2f cannot be divided by 0", e.dividend)
}
func divider(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("no dividing by 0") // rather than creating a custom error, we can errors.New() method.
		//return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func testDivider(dividend, divisor float64) {
	defer fmt.Println("=======================")
	fmt.Printf("Dividing %.2f by %.2f ...\n", dividend, divisor)

	res, err := divider(dividend, divisor)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
