package medianoftwosortedarrays

import (
	"testing"
)

func TestMedianOfTwoSortedArrays(t *testing.T) {
	testCases := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected float64
	}{
		{"test1", []int{1, 3}, []int{2}, 2.0},
		{"test2", []int{1, 2}, []int{3, 4}, 2.5},
		{"test3", []int{1, 2, 4}, []int{3}, 2.5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := findMedianSortedArrays(tc.nums1, tc.nums2)

			if output != tc.expected {
				t.Errorf("For nums1=%v and nums2=%v, expected %v, but got %v",
					tc.nums1, tc.nums2, tc.expected, output)
			}
		})
	}
}
