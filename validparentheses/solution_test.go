package validparentheses

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"test1", "()", true},
		{"test2", "()[]{}", true},
		{"test3", "(]", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := isValid(tc.input)

			if output != tc.expected {
				t.Errorf("For input=%v, expected %v, but got %v",
					tc.input, tc.expected, output)
			}
		})
	}
}
