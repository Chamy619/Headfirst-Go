package main

import (
	"fmt"
	"headfirst-go/chapter08/magazine"
)

func main() {
	var s magazine.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)
}