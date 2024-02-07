package one

import (
	"testing"
)

func TestProcessLine(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
		{"eightknfssevenfive6jnpklczrpfeightwol", 82},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := processLine(test.input)

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if result != test.expected {
				t.Errorf("Expected %d, got %d, input '%s'", test.expected, result, test.input)
			}
		})
	}
}
