package main

import (
	"fmt"
	"headfirst-go/chapter10/calendar"
	"log"
)


func main() {
	date := calendar.Date{}
	err := date.SetYear(2024)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(2)
	if err != nil {
		log.Fatal(err)
	}
	
	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}