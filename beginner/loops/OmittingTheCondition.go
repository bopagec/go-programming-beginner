package main

import "fmt"

// Loops in Go can omit sections of a for loop. For example, the CONDITION (middle part) can be omitted which causes the loop to run forever.
/*
for INITIAL; ; AFTER {
  // do something forever
}
*/

/*
ASSIGNMENT
Complete the maxMessages function. Given a cost threshold, it should calculate the maximum number of messages that can be sent.

Each message costs 1.0, plus an additional fee. The fee structure is:

1st message: 1.0 + 0.00
2nd message: 1.0 + 0.01
3rd message: 1.0 + 0.02
4th message: 1.0 + 0.03
*/
func main() {
	test2(10)
	test2(20)
	test2(30)
	test2(40)
	test2(50)
}

func maxMessages(thresh float64) int {
	total := 0.0
	for i := 0; ; i++ {
		total += 1 + 0.01*float64(i)
		if total >= thresh {
			return i
		}
	}
}
func test2(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	max := maxMessages(thresh)
	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
	fmt.Println("===============================================================")
}
