package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	
	score, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		log.Fatal(err)
	}
	if  score >= 60 {
		fmt.Println("passing")
	} else {
		fmt.Println("failing")
	}
}