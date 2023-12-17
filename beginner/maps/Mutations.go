package main

import (
	"errors"
	"fmt"
	"sort"
)

// documentation
// https://go.dev/blog/maps
/*
MUTATIONS
INSERT AN ELEMENT
m[key] = elem

GET AN ELEMENT
elem = m[key]

DELETE AN ELEMENT
delete(m, key) // The delete function doesn’t return anything, and will do nothing if the specified key doesn’t exist.

CHECK IF A KEY EXISTS
elem, ok := m[key]

To test for a key without retrieving the value, use an underscore in place of the first value:

_, ok := m["route"]
If key is in m, then ok is true. If not, ok is false.

If key is not in the map, then elem is the zero value for the map's element type.

To iterate over the contents of a map, use the range keyword:

for key, value := range m {
    fmt.Println("Key:", key, "Value:", value)
}

ASSIGNMENT
It's important to keep up with privacy regulations and to respect our user's data. We need a function that will delete user records.

Complete the deleteIfNecessary function.

Check the scheduledForDeletion bool to determine if they are scheduled for deletion or not.

If the user doesn't exist in the map, return the error not found.
If they exist but aren't scheduled for deletion, return deleted as false with no errors.
If they exist and are scheduled for deletion, return deleted as true with no errors and delete their record from the map.

NOTE ON PASSING MAPS
Like slices, maps are also passed by reference into functions.
This means that when a map is passed into a function we write, we can make changes to the original, we don't have a copy.
*/
func deleteIfNecessary(customers map[string]customer, name string) (deleted bool, err error) {
	cust, ok := customers[name]

	if !ok {
		return false, errors.New("customer not found")
	}

	if cust.scheduledForDeletion {
		delete(customers, name)
		return true, nil
	}
	return false, nil
}

// don't touch below this line

type customer struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func test2(customers map[string]customer, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("====================================")
	deleted, err := deleteIfNecessary(customers, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if deleted {
		fmt.Println("Deleted:", name)
		return
	}
	fmt.Println("Did not delete:", name)
}

func main() {
	customers := map[string]customer{
		"john": {
			name:                 "john",
			number:               18965554631,
			scheduledForDeletion: true,
		},
		"elon": {
			name:                 "elon",
			number:               19875556452,
			scheduledForDeletion: true,
		},
		"breanna": {
			name:                 "breanna",
			number:               98575554231,
			scheduledForDeletion: false,
		},
		"kade": {
			name:                 "kade",
			number:               10765557221,
			scheduledForDeletion: false,
		},
	}
	test2(customers, "john")
	test2(customers, "musk")
	test2(customers, "santa")
	test2(customers, "kade")

	keys := []string{}
	for name := range customers {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	fmt.Println("Final map keys:")
	for _, name := range keys {
		fmt.Println(" - ", name)
	}
}
