package datafile

import (
	"bufio"
	"os"
)

func GetStrings(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return lines, nil
}