package longestcommonprefix

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected string
	}{
		{"test1", []string{"flower", "flow", "flight"}, "fl"},
		{"test2", []string{"dog", "racecar", "car"}, ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := longestCommonPrefix(tc.input)

			if output != tc.expected {
				t.Errorf("For input=%v, expected %v, but got %v",
					tc.input, tc.expected, output)
			}
		})
	}
}
