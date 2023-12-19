package main

import (
	"errors"
	"fmt"
)

/*
ASSIGNMENT
The Mailio API needs a very robust error-logging system so we can see when things are going awry in the back-end system.
We need a function that can create a custom "logger" (a function that prints to the console) given a specific formatter.

Complete the getLogger function.
It should return a new function that prints the formatted inputs using the given formatter function.
The inputs should be passed into the formatter function in the order they are given to the logger function.
*/

// getLogger takes a function that formats two strings into a single string and returns a function that
// formats two strings but prints the result instead of returning it
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(s string, s2 string) {
		fmt.Println(formatter(s, s2))
	}
}

func test2(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("====================================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func colonDelimit(first, second string) string {
	return first + ": " + second
}
func commaDelimit(first, second string) string {
	return first + ", " + second
}

func main() {
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	test2("Error on database server", dbErrors, colonDelimit)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}
	test2("Error on mail server", mailErrors, commaDelimit)
}
