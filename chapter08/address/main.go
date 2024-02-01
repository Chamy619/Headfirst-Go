package main

import (
	"fmt"
	"headfirst-go/chapter08/magazine"
)

func main() {
	address := magazine.Address{Street: "123 Oak St", City: "Omaha", State: "Ne", PostalCode: "68111"}
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Address = address
	fmt.Println(subscriber.Address)
	fmt.Println(subscriber.Address.City)
	fmt.Println(subscriber.City)
}