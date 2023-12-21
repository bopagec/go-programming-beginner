package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	user1 := user{
		name: "tyrone",
		age:  9,
	}

	user1.change()

	fmt.Printf("user after changed, %v", user1)

}

func (u *user) change() {
	u.age = u.age + 1
	u.name = "Tyrone Bopage"
}
