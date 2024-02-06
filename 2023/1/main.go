package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func processLine(line string) (int, error) {
	// In Go it is common to define multiple variables at once
	var firstDigit, lastDigit int

	// Here we are iterating through each character of the line
	for _, char := range line {
		if unicode.IsNumber(char) {
			// In Go, we can convert a rune to an int by subtracting the rune '0'
			intChar := int(char - '0')

			if firstDigit == 0 {
				firstDigit = intChar
			}

			lastDigit = intChar
		} else if unicode.IsLetter(char) {
			continue
		} else {
			log.Fatal("Unknown character")
		}

	}

	concatenatedDigits, err := strconv.Atoi(strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit))

	if err != nil {
		fmt.Println(err)
	}

	return concatenatedDigits, nil

}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Scheduling now to close the file when the function is done
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	// Iterate through each line of the file allowing us to efficiently handle each item
	for scanner.Scan() {
		lineSum, err := processLine(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		sum += lineSum
	}

	// In Go it is common to define an error variable and check it immediately
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
