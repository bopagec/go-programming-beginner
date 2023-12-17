package main

import "fmt"

/*
RANGE
Go provides syntactic sugar to iterate easily over elements of a slice:
*/
/*
ASSIGNMENT
We need to be able to quickly detect bad words in the messages our system sends.

Complete the indexOfFirstBadWord function.
If it finds any bad words in the message it should return the index of the first bad word in the msg slice.
This will help us filter out naughty words from our messaging system.
If no bad words are found, return -1 instead.

For example:

fruits := []string{"apple", "banana", "grape"}
for i, fruit := range fruits {
    fmt.Println(i, fruit)
}
// 0 apple
// 1 banana
// 2 grape
Use the range keyword.
*/
func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, message := range msg {
		for _, badWord := range badWords {
			if message == badWord {
				return i
			}
		}
	}
	return -1
}

// don't touch below this line

func test6(msg []string, badWords []string) {
	i := indexOfFirstBadWord(msg, badWords)
	fmt.Printf("Scanning message: %v for bad words:\n", msg)
	for _, badWord := range badWords {
		fmt.Println(` -`, badWord)
	}
	fmt.Printf("Index: %v\n", i)
	fmt.Println("====================================")
}

func main() {
	badWords := []string{"crap", "shoot", "dang", "frick"}
	message := []string{"hey", "there", "john"}
	test6(message, badWords)

	message = []string{"ugh", "oh", "my", "frick"}
	test6(message, badWords)

	message = []string{"what", "the", "shoot", "I", "hate", "that", "crap"}
	test6(message, badWords)
}
