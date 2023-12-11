package main

import "fmt"

// a function can return a value that the caller doesn't care about we can explicitly ignore variables by using the underscore: _
// in Go, you cannot have unused variables in your program
// because Go does not allow you to have unused variables, Go will allow you to ignore some variables.
func main() {
	firstName, lastName := getNames()
	// if you declare here only the firstName you will get a compiler error as lastName is not used.
	fmt.Println("your full name, ", firstName, lastName)

	fName, _ := getNames()
	fmt.Println("your first name, ", fName)
}

// when you have multiple return values we should put them inside parenthesis
func getNames() (string, string) {
	return "Tyrone", "Bopage"
}
