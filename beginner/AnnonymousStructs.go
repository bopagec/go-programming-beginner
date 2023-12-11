package main

import "fmt"

func main() {
	// notice there is no type for myCar stuct
	myCar := struct {
		make  string
		model string

		// wheel is a field containing anonymous struct
		// notice there is no type for wheel struct
		wheel struct {
			radius   int
			meterial string
		}
	}{
		make:  "BMW",
		model: "5 Series",
	}

	fmt.Printf("my car make: %s , model:%s", myCar.make, myCar.model)
}
