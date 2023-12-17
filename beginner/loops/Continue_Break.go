package main

import (
	"fmt"
)

// continue and break statements in for loop similar to java
// break will exit the loop while continue will continue to the next iteration without further actions.
/*
ASSIGNMENT
As an easter egg, we decided to reward our users with a free text message if they send a prime number of text messages this year.

Complete the printPrimes function. It should print all of the prime numbers up to and including max. It should skip any numbers that are not prime.


BREAKDOWN
We skip even numbers because they can't be prime
We only check up to the square root of n because anything higher than the square root has no chance of multiplying evenly into n
We start checking at 2 because 1 is not prime
*/
func main() {
	test4(10)
	test4(20)
	test4(30)
}

func test4(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func printPrimes(max int) {
	for i := 2; i <= max; i++ {
		if i == 2 {
			fmt.Println(i)
			continue
		}
		if i%2 == 0 {
			continue
		}

		isPrime := true
		for x := 3; x*x <= i; x++ {
			if i%x == 0 {
				isPrime = false
				break
			}
		}

		if !isPrime {
			continue
		}

		fmt.Println(i)
	}
}
