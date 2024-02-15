package main

import (
	"fmt"
	"headfirst-go/chapter14/prose"
)

func main() {
	phrases := []string{"my parents", "a radeo clown"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
}