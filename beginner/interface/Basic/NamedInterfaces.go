package main

func main() {

}

// Let's add some named arguments adn return data to make it more readable and clear
// Much better, we can see what the expectation are now from the interface.
type Copier interface {
	// named interface method
	Copy(sourceFile string, destinationFile string) (bytesCopied int)

	// unnamed interface method
	diff(string, string) int
}
