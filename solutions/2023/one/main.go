package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/janschloss/adventofcode/utils"
)

func main() {
	utils.ProcessInput("input.txt", processLine)
}

// I wanted to try goroutines here, even though my benchmarks show that this is
// actually slower than the non-goroutine version, most likely due to the overhead
// I'm leaving it here as an example of how to use goroutines
func processLine(line string) (int, error) {
	firstDigitChan := make(chan int)
	lastDigitChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// Find first digit from the left
	// We use line[i:] to get the substring from i to the end of the string
	go func() {
		defer wg.Done()
		for i := 0; i < len(line); i++ {
			if digit, ok := findDigit(line[i:]); ok {
				firstDigitChan <- digit
				return
			}
		}
	}()

	// Find last digit from the right
	go func() {
		defer wg.Done()
		for i := len(line) - 1; i >= 0; i-- {
			if digit, ok := findDigit(line[i:]); ok {
				lastDigitChan <- digit
				return
			}
		}
	}()

	// This is how you wait for multiple goroutines to finish
	go func() {
		wg.Wait()
		close(firstDigitChan)
		close(lastDigitChan)
	}()

	firstDigit, lastDigit := 0, 0

	for i := 0; i < 2; i++ {
		select {
		case digit := <-firstDigitChan:
			firstDigit = digit
		case digit := <-lastDigitChan:
			lastDigit = digit
		}
	}

	// 0 is not a valid digit in this challenge
	if firstDigit == 0 || lastDigit == 0 {
		return 0, fmt.Errorf("no digit found")
	}

	// Here we concatenate the digits which is required by the problem
	return firstDigit*10 + lastDigit, nil

}

// We define this outside of the function to avoid creating it every time the function is called
var wordToDigit = map[string]int{
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
