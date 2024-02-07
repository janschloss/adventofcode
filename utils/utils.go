package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// In Go a capitalized function name means it is public
func ProcessInput(filePath string, processor func(string) (int, error)) {
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
		lineSum, err := processor(scanner.Text())

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
