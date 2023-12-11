package main

import "fmt"

func main() {
	r := rect{width: 10, height: 5}
	fmt.Printf("area of the rect: %v\n", r.area()) // 50

	info := authenticationInfo{username: "tyrone", password: "tau"}
	fmt.Printf(info.getBasicAuth())
}

type rect struct {
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

type authenticationInfo struct {
	username string
	password string
}

func (info authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", info.username, info.password)
}
