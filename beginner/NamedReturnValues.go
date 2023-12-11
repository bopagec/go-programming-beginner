package main

import "fmt"

func main() {
	yearsUntilDrive, yearsUntilAdult, yearsUntilDrinking := yearsUntilEvents(9)

	fmt.Println("yearsUntilDrive:", yearsUntilDrive)
	fmt.Println("yearsUntilAdult:", yearsUntilAdult)
	fmt.Println("yearsUntilDrinking", yearsUntilDrinking)
	x, y := getCoordImpl()
	fmt.Println("implicit returns: ", x, y)
}

// named variables with implicit return
func yearsUntilEvents(age int) (
	yearsUntilDrive int,
	yearsUntilAdult int,
	yearsUntilDrinking int) {

	yearsUntilDrive = 16 - age
	if yearsUntilDrive < 0 {
		yearsUntilDrive = 0
	}

	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}

	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}

	return // implicit return
	//return yearsUntilDrive,yearsUntilAdult,yearsUntilDrinking // explicit return which is recommended.
}

// implicit return
// we're using names for returns as x and y
func getCoordImpl() (x int, y int) {
	return // this will return the default value of x and y
}

// explicit return ( recommended )
// we're using names for returns as x and y
func getCoordExpl() (x int, y int) {
	return x, y // this will return the default value of x and y explicitly
}
