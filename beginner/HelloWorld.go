package main

import "fmt"

func main() {
	var first_name = "Tyrone "
	var last_name = "Bopage"

	isAddress := true // or var isAddress = true
	houseNo := 14
	roadName := "Boxhill"
	postCode := "HP2 5YR"

	price := 4.56

	mileage, company := 87000, "Tesla"

	accountAge := 2.6

	const premiumPlanName = "Premium Plan" // constant
	//premiumPlanName = "test" // cannot change

	fmt.Println("Hello World", first_name+" "+last_name)
	fmt.Println(isAddress, houseNo, roadName, postCode)

	fmt.Printf("The type of houseNo, %T\n", houseNo) // %T = type of the variable, \n = new line
	fmt.Printf("The type of price, %T\n", price)     // %T = type of the variable, \n = new line
	fmt.Println(mileage, company)

	fmt.Println("Accont Age in", int(accountAge), "years") //conversion of flaot to int
	fmt.Println("House no in", float64(houseNo), "float")  //conversion of int to float64
	fmt.Println(premiumPlanName)                           //conversion of int to float64

	fmt.Printf("I am %v years old\n", 45)             // %v is the default for everything if you not sure of the data type you can use this
	fmt.Printf("I am %s years old\n", "way too many") // interpolate a string
	fmt.Printf("I am %d years old\n", 10)             // interpolate a integer in decimal form
	fmt.Printf("I am %f years old\n", 45.1)           // interpolate a decimal (float)
	fmt.Printf("I am %.2f years old\n", 45.1456)      // rounds the number to 2 decimal places

	// SprintF will return the formatted string

	const name = "Tyrone Bopage"
	const openRate = 30.5
	msg := fmt.Sprintf("Hi %s, your open rate is %.2f percent", name, openRate)

	fmt.Printf(msg)

}
