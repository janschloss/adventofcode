package main

import (
	"testing"
)

// Assume processLine is the function you want to benchmark
// Results without Goroutines:  1076	   1112701 ns/op	   29388 B/op	    1005 allocs/op
// Result with Goroutines:  	282	   	   4171671 ns/op	   365889 B/op	    7008 allocs/op
func BenchmarkMain(b *testing.B) {
	// The b.N value is chosen dynamically by the benchmarking framework to obtain meaningful results
	for i := 0; i < b.N; i++ {
		main()
	}
}

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
