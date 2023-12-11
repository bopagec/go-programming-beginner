package main

import (
	"errors"
	"fmt"
)

// Guard Clause are much popular in modern programming
func main() {
	res, er := divide(10, 5)
	fmt.Println("divide test ", res, er)
}

func divide(divident int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't devide by zero")
	}

	return divident / divisor, nil
}
