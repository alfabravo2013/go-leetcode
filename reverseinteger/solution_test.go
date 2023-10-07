package reverseinteger

import "testing"

func TestReverse(t *testing.T) {
	testCases := []struct {
		name     string
		x        int
		expected int
	}{
		{"test1", 123, 321},
		{"test2", 1234567, 7654321},
		{"test3", 120, 21},
		{"test4", 1534236469, 0},
		{"test5", -123, -321},
		{"test6", 102, 201},
		{"test7", 1463847412, 2147483641},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := reverse(tc.x)

			if output != tc.expected {
				t.Errorf("For x=%v, expected %v, but got %v",
					tc.x, tc.expected, output)
			}
		})
	}
}
