package main

import (
	"errors"
	"fmt"
)

/*
This variable m is a map of string keys to int values:

map key can only be comparable types such as int, bool structs etc ( structs are comparable )
mpa key cannot have non-comparable items such as maps or slice

var m map[string]int
Map types are reference types, like pointers or slices, and so the value of m above is nil; it doesn’t point to an initialized map.
A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don’t do that.
To initialize a map, use the built in make function:

m = make(map[string]int)
The make function allocates and initializes a hash map data structure and returns a map value that points to it.
The specifics of that data structure are an implementation detail of the runtime and are not specified by the language itself.
In this article we will focus on the use of maps, not their implementation.

MAPS
Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping.

The zero value of a map is nil.

We can create a map by using a literal or by using the make() function:

ages := make(map[string]int) // map names to the age (key:value)
ages["John"] = 37
ages["Mary"] = 24
ages["Mary"] = 21 // overwrites 24
Copy icon

	ages = map[string]int{
	  "John": 37,
	  "Mary": 21,
	}

The len() function works on a map, it returns the total number of key/value pairs.

	ages = map[string]int{
	  "John": 37,
	  "Mary": 21,
	}

MAP LITERALS
Maps can be constructed using the usual composite literal syntax with colon-separated key-value pairs, so it's easy to build them during initialization.

	var timeZone = map[string]int{
	    "UTC":  0*60*60,
	    "EST": -5*60*60,
	    "CST": -6*60*60,
	    "MST": -7*60*60,
	    "PST": -8*60*60,
	}

fmt.Println(len(ages)) // 2

ASSIGNMENT
We can speed up our contact-info lookups by using a map! Looking up a value in a map by its key is much faster than searching through a slice.

Complete the getUserMap function. It takes a slice of names and a slice of phone numbers, and returns a map of name -> user structs and potentially an error.
A user struct just contains a user's name and phone number.

If the length of names and phoneNumbers is not equal, return an error with the string "invalid sizes".

The first name in the names slice matches the first phone number, and so on.
*/
func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)

	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	for i, name := range names {
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumbers[i],
		}
	}
	return userMap, nil
}

// don't touch below this line

type user struct {
	name        string
	phoneNumber int
}

func test(names []string, phoneNumbers []int) {
	fmt.Println("Creating map...")
	defer fmt.Println("====================================")
	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value:\n", name)
		fmt.Println(" - name:", users[name].name)
		fmt.Println(" - number:", users[name].phoneNumber)
	}
}

func main() {
	test(
		[]string{"John", "Bob", "Jill"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"John", "Bob"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"George", "Sally", "Rich", "Sue"},
		[]int{20955559812, 38385550982, 48265554567, 16045559873},
	)
}
