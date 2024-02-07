package one

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Scheduling now to close the file when the function is done
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	// Iterating through each line of the file
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

func processLine(line string) (int, error) {
	// In Go it is common to define multiple variables at once
	var firstDigit, lastDigit int

	// Find first digit from the left
	// We use line[i:] to get the substring from i to the end of the string
	for i := 0; i < len(line); i++ {
		if digit, ok := findDigit(line[i:]); ok {
			firstDigit = digit
			break
		}
	}

	// Find last digit from the right
	for i := len(line) - 1; i >= 0; i-- {
		if digit, ok := findDigit(line[i:]); ok {
			lastDigit = digit
			break
		}
	}

	// Here we concatenate the digits which is required by the problem
	return firstDigit*10 + lastDigit, nil

}

// We define this outside of the function to avoid creating it every time the function is called
var wordToDigit = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func findDigit(partialLine string) (int, bool) {
	// Check for a digit as the first character
	firstChar := partialLine[0]

	if firstChar >= '0' && firstChar <= '9' {
		return int(firstChar - '0'), true
	}

	// Check for word representations of digits.
	for word, digit := range wordToDigit {
		if strings.HasPrefix(partialLine, word) {
			return digit, true
		}
	}

	return 0, false
}
