package problems

import (
	"testing"
)

func TestIsPalindrom(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{"example 1", 121, true},
		{"example 2", -121, false},
		{"example 3", 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.x)
			if got != tt.want {
				t.Errorf("isPalindrom(%v) = %v; want %v", tt.x, got, tt.want)
			}
		})
	}
}
