package main

import "testing"

func TestReversePolishNotation(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Simple operations
		{"3 4 +", "7"},
		{"10 4 -", "6"},
		{"3 4 *", "12"},
		{"12 3 /", "4"},
		// A more complex expression: 5 + ((1 + 2) * 4) - 3 = 14
		{"5 1 2 + 4 * + 3 -", "14"},
		// Edge cases
		{"", "0"},   // empty input, no operations performed
		{"4 +", "0"}, // not enough operands; this will print "error" and return 0
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ReversePolishNotation(tt.input)
			if result != tt.expected {
				t.Errorf("For input %q, expected %q but got %q", tt.input, tt.expected, result)
			}
		})
	}
}
