package romantointeger

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"test1", "III", 3},
		{"test2", "LVIII", 58},
		{"test3", "MCMXCIV", 1994},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := romanToInt(tc.input)

			if output != tc.expected {
				t.Errorf("For input=%v, expected %v, but got %v",
					tc.input, tc.expected, output)
			}
		})
	}
}
