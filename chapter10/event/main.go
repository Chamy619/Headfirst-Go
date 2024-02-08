package main

import (
	"fmt"
	"headfirst-go/chapter10/calendar"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetYear(2024)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetMonth(2)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}

	// err = event.SetTitle("An extremely long and impractical title")
	err = event.SetTitle("Mom's birthday")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}