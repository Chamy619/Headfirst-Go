package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// guess 프로그램은 플레이어가 난수를 맞히는 게임
func main() {
	restGuessCount := 10
	target := rand.Intn(99) + 1
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	for restGuessCount > 0 {
		fmt.Printf("You have %d guesses left.\n", restGuessCount)
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		guess, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			log.Fatal(err)
		}

		if guess == target {
			fmt.Println("Good job! You guessed it!")
			break;
		} else if guess > target {
			fmt.Println("Oops. Your guess was High.")
		} else {
			fmt.Println("Oops. Your guess was Low.")
		}

		restGuessCount--
	}

	if restGuessCount == 0 {
		fmt.Printf("You didn't guess my number. It was: [%d]\n", target)
	}
}