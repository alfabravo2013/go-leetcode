package integertoroman

import (
	"testing"
)

func TestIntegerToRoman(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected string
	}{
		{"test1", 3, "III"},
		{"test2", 58, "LVIII"},
		{"test3", 1994, "MCMXCIV"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := intToRoman(tc.input)

			if output != tc.expected {
				t.Errorf("For input=%v, expected %v, but got %v",
					tc.input, tc.expected, output)
			}
		})
	}
}
