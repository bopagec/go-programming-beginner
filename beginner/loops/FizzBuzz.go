package main

import "fmt"

/*
ASSIGNMENT
We're hiring engineers at Textio, so time to brush up on the classic "Fizzbuzz" game,
a coding exercise that has been dramatically overused in coding interviews across the world.

Complete the fizzbuzz function that prints the numbers 1 to 100 inclusive each on their own line,
but substitutes multiples of 3 for the text fizz and multiples of 5 for buzz.
For multiples of 3 AND 5 print instead fizzbuzz.
*/
func main() {
	fizzbuzz()
}

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		fizz := i%3 == 0
		buzz := i%5 == 0

		if fizz && buzz {
			fmt.Printf("fizzbuzz (%d)\n", i)
		} else if fizz {
			fmt.Printf("fizz (%d)\n", i)
		} else if buzz {
			fmt.Printf("buzz (%d)\n", i)
		} else {
			fmt.Println(i)
		}
	}
}
