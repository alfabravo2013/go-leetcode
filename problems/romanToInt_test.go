package problems

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"example 1", "III", 3},
		{"example 2", "LVIII", 58},
		{"example 3", "MCMXCIV", 1994},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.s)
			if got != tt.want {
				t.Errorf("romanToInt(%v) = %v; want %v", tt.s, got, tt.want)
			}
		})
	}
}
