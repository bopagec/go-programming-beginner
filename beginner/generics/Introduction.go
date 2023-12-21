package main

import "fmt"

/*
Why Generics ?
1) keep the code compile time safe
2) code re-usability (DRY) Don't Repeat Yourself
3) avoid runtime checking


GENERICS IN GO
As we've mentioned, Go does not support classes. For a long time, that meant that Go code couldn't easily be reused in many circumstances. For example, imagine some code that splits a slice into 2 equal parts. The code that splits the slice doesn't really care about the values stored in the slice. Unfortunately in Go we would need to write it multiple times for each type, which is a very un-DRY thing to do.

func splitIntSlice(s []int) ([]int, []int) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
Copy icon
func splitStringSlice(s []string) ([]string, []string) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
Copy icon
In Go 1.18 however, support for generics was released, effectively solving this problem!

TYPE PARAMETERS
Put simply, generics allow us to use variables to refer to specific types.
This is an amazing feature because it allows us to write abstract functions that drastically reduce code duplication.
*/

func splitAnySplice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2

	return s[:mid], s[mid:]
}

func main() {
	first, second := splitAnySplice([]int{0, 1, 2, 3, 4})
	fmt.Println(first, second)
}
