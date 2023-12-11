package main

import "fmt"

/*
Go, is not an object-oriented language. However, embedded structs provide
a kind of data-only inheritance that can be useful at times, Keep in mind,
Go does not support classes or inheritance in the complete sense,
Embedded structs are just a way to elevate and share fields between struct definitions.
*/
func main() {
	type car struct {
		make  string
		model string
	}

	type truck struct {
		car
		bedsize int
	}

	lanesTruct := truck{
		bedsize: 10,
		car: car{
			make:  "toyota",
			model: "camry",
		},
	}

	fmt.Printf("truck details, bedsize: %d, make: %s, model: %s", lanesTruct.bedsize, lanesTruct.make, lanesTruct.model)
}
