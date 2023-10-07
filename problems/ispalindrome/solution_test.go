package ispalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		x        int
		expected bool
	}{
		{"test1", 121, true},
		{"test2", -121, false},
		{"test3", 10, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := isPalindrome(tc.x)

			if output != tc.expected {
				t.Errorf("For x=%v, expected %v, but got %v",
					tc.x, tc.expected, output)
			}
		})
	}
}

func TestIsPalindromeEfficient(t *testing.T) {
	testCases := []struct {
		name     string
		x        int
		expected bool
	}{
		{"test1", 121, true},
		{"test2", -121, false},
		{"test3", 10, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := isPalindromeEfficient(tc.x)

			if output != tc.expected {
				t.Errorf("For x=%v, expected %v, but got %v",
					tc.x, tc.expected, output)
			}
		})
	}
}
