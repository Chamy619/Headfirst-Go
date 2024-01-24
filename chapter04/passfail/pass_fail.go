package main

import (
	"fmt"
	"headfirst-go/chapter04/keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	
	if grade >= 60 {
		fmt.Println("A grade of", grade, "is passing")
	} else {
		fmt.Println("A grade of", grade, "is failing")
	}
}