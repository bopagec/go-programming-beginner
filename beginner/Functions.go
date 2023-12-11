package main

import "fmt"

func main() {
	fmt.Println(sub(10, 2))
	fmt.Println("Tyrone", "Bopage")
	fmt.Println(add(10, 20))
}

func sub(x int, y int) int {
	return x - y
}

func concat(s1 string, s2 string) string {
	return s1 + " " + s2
}

func add(x, y int) int {
	return x + y
}
