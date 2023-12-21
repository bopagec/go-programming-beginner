package main

/*
ASSIGNMENT
At Mailio we often store all the emails of a given email campaign in memory as a slice.
We store payments for a single user in the same way.

Complete the getLast() function.
It should be a generic function that returns the last element from a slice,
no matter the types stored in the slice. If the slice is empty, it should return the zero value of the type.

TIP: ZERO VALUE OF A TYPE
Creating a variable that's the zero value of a type is easy:

var myZeroInt int

It's the same with generics, we just have a variable that represents the type:

var myZero T
*/
import "fmt"

type email struct {
	message       string
	senderEmail   string
	receiverEmail string
}

type payment struct {
	amount        int
	senderEmail   string
	receiverEmail string
}

func main() {
	test([]email{}, "email")
	test([]email{
		{
			"Hi Margo",
			"janet@example.com",
			"margo@example.com",
		},
		{
			"Hey Margo I really wanna chat",
			"janet@example.com",
			"margo@example.com",
		},
		{
			"ANSWER ME",
			"janet@example.com",
			"margo@example.com",
		},
	}, "email")

	test([]payment{
		{
			5,
			"jane@example.com",
			"sally@example.com",
		},
		{
			25,
			"jane@example.com",
			"mark@example.com",
		},
		{
			1,
			"jane@example.com",
			"sally@example.com",
		},
		{
			16,
			"jane@example.com",
			"margo@example.com",
		},
	}, "payment")
}

func getLast[T any](s []T) (t T) {
	if len(s) == 0 {
		var zeroVal T
		return zeroVal
	}
	return s[len(s)-1]
}
func test[T any](s []T, desc string) {
	last := getLast(s)
	fmt.Printf("Getting last %v from slice of length: %v\n", desc, len(s))

	for i, v := range s {
		fmt.Printf("Item #%v: %v\n", i+1, v)
	}
	fmt.Printf("Last item in list: %v\n", last)
	fmt.Println(" --- ")
}
