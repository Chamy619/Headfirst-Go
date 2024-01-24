// Package keyboard reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat reads a floating-point number from the keyboards.
// It returns the number read and any error encountered.
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	number, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}