package medianoftwosortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, 0)

	// todo: try using two pointers to find the median

	i1 := 0
	i2 := 0

	for i1 < len(nums1) && i2 < len(nums2) {
		if nums1[i1] <= nums2[i2] {
			merged = append(merged, nums1[i1])
			i1++
		} else {
			merged = append(merged, nums2[i2])
			i2++
		}
	}

	for i1 < len(nums1) {
		merged = append(merged, nums1[i1])
		i1++
	}

	for i2 < len(nums2) {
		merged = append(merged, nums2[i2])
		i2++
	}

	total := len(merged)
	if total%2 == 0 {
		med1 := merged[total/2-1]
		med2 := merged[total/2]
		return (float64(med1) + float64(med2)) / 2.0
	}
	return float64(merged[total/2])
}
