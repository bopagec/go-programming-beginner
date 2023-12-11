package main

import "fmt"

// interfaces are implemented implicitly in Go,
// there is no keyword to implement interfaces in Go.
// a type can implement any number of interfaces in Go.
func main() {
	testEmployees(contractor{
		employee:     employee{name: "Tyrone Bopage"},
		hourlyPay:    50,
		hoursPerYear: 1000,
	})

	testEmployees(permanent{
		employee: employee{name: "Nush Wijayakoon"},
		salary:   45000,
	})
}

type employee struct {
	name string
}

func testEmployees(emp employeeInt) {
	fmt.Println(emp.getName())
	fmt.Println(emp.getSalary())
	fmt.Println("============================")
}

func (e employee) getName() string {
	return e.name
}

type employeeInt interface {
	getSalary() int
	getName() string
}

type contractor struct {
	employee
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getSalary() int {
	return c.hoursPerYear * c.hourlyPay
}

type permanent struct {
	employee
	salary int
}

func (p permanent) getSalary() int {
	return p.salary
}
