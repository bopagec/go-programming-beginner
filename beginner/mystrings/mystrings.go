package mystrings

// Notice that we need to capitalize the first letter in the method name to expose the method in the package
// if not, then wwe won't be able to access this method in mystrings package.
func Reverse(s string) string {
	result := ""
	for _, c := range s {
		result = string(c) + result
	}
	return result
}
