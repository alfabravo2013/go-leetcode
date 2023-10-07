package twosum

import (
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{"test1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"test2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"test3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := twosum(tc.nums, tc.target)

			if !testEq(output, tc.expected) {
				t.Errorf("For nums=%v and target=%v, expected %v, but got %v",
					tc.nums, tc.target, tc.expected, output)
			}
		})
	}
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
