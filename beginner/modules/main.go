package main

import (
	"fmt"
	tinytime "github.com/wagslane/go-tinytime"
	"time"
)

func main() {
	tt := tinytime.New(1585750347)
	tt = tt.Add(time.Hour * 48)
	fmt.Println(tt)
}

// we use go get to download the code
// go get github.com/wagslane/go-tinytime
// this will automatically add it to go.mod file
// go build
// go run main.go
