package main

/*
INTRODUCTION TO POINTERS
Pointers hold the memory address of a value.

As we have learned, a variable is a named location in memory that stores a value.
We can manipulate the value of a variable by assigning a new value to it or by performing operations on it.
When we assign a value to a variable, we are storing that value in a specific location in memory.

x := 42
y: = x

"x" is the name of a location in memory. That location is storing the integer value of 42
"y" is the name of a location in memory that hold it's own value of 42
x and y have 2 different memory addresses and, if y change it will not going to change the value of x.

name addr  value
x    169   42
y    170   42

A POINTER IS A VARIABLE
A pointer is a variable that stores the memory address of another variable.
This means that a pointer "points to" the location of where the data is stored NOT the actual data itself.

The & operator generates a pointer to its operand.

myStr := "hello"
myStrPtr := &myString

name 		addr	value
myStr		171		hello
myStrPnt	172		171 ( location of the memory address)

A pointer's zero value is nil

The * syntax defines a pointer:
POINTER is a specific type in Golang
var p *int
var q *string

The * dereferences a pointer to gain access to the value

fmt.Println(*myStrPtr) // read myStr through the pointer
*myStrPtr = "world"    // set myStr through the pointer

Unlike C, Go has no pointer arithmetic

WHY ARE POINTERS USEFUL?
Pointers allow us to manipulate data in memory directly, without making copies or duplicating data.
This can make programs more efficient and allow us to do things that would be difficult or impossible without them.

JUST BECAUSE YOU CAN DOESN'T MEAN YOU SHOULD
We're doing this exercise to understand that pointers can be used in this way.
That said, pointers can be very dangerous.
It's generally a better idea to have your functions accept non-pointers and return new values rather than mutating pointer inputs.

remember, * is used in 2 different places
1) dereference a pointer ie *myStr
2) use pointer types ie *int

& is used to create new references to the pointer
*/
