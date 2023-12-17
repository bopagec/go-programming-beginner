package main

import "fmt"

// The built-in append function is used to dynamically add elements to a slice:
// If the underlying array is not large enough, append() will create a new underlying array and point the slice to it.

// func append(slice []Type, elems ...Type) []Type
// Notice that append() is variadic, the following are all valid:
// slice = append(slice, oneThing)
// slice = append(slice, firstThing, secondThing)
// slice = append(slice, anotherSlice...)

/*
ASSIGNMENT
We've been asked to "bucket" costs per day, in a given period.
Complete the getCostsByDay function. It should return a slice of float64s, where each element is the total cost for that day.
The length of the slice should be equal to the number of days represented in the costs slice up to the last day represented in the slice.

Only one "bucket" of costs per day, and include a "bucket" for days without any costs.

Days are numbered and start at 0.
*/
func main() {
	test4([]cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{2, 2.5},
		{3, 3.6},
		{3, 2.7},
		{4, 3.34},
	})
	test4([]cost{
		{0, 1.0},
		{10, 2.0},
		{3, 3.1},
		{2, 2.5},
		{1, 3.6},
		{2, 2.7},
		{4, 56.34},
		{13, 2.34},
		{28, 1.34},
		{25, 2.34},
		{30, 4.34},
	})
}

type cost struct {
	day   int
	value float64
}

func test4(costs []cost) {
	fmt.Printf("Creating daily buckets for %v costs...\n", len(costs))
	costsByDay := getCostsByDay(costs)
	fmt.Println("Costs by day:")
	for i := 0; i < len(costsByDay); i++ {
		fmt.Printf(" - Day %v: %.2f\n", i, costsByDay[i])
	}
	fmt.Println("===== END REPORT =====")
}

func getCostsByDay(costs []cost) []float64 {
	costByDay := []float64{} // similar to make([] float64,0)

	for i := 0; i < len(costs); i++ {
		cost := costs[i]

		for cost.day >= len(costByDay) {
			costByDay = append(costByDay, 0.0)
		}
		costByDay[cost.day] += cost.value

	}

	return costByDay
}
