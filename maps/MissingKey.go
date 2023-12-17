package main

import "fmt"

func main() {
	attended := map[string]bool{
		"Ann": true,
		"Joe": true,
	}

	if attended["tyrone"] {
		fmt.Println("tyrone was in the meeting")
	} else {
		fmt.Println("tyrone was not in the meeting")
	}

	i, ok := attended["tyrone"] // missing key will be false for bool, missing key will be 0 for int

	fmt.Printf("i: %v ok: %v", i, ok) // i: false ok: false
}
