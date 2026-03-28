package problems

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for idx, n := range nums {
		delta := target - n
		dix, ok := m[n]
		if ok {
			if n+delta == target {
				return []int{idx, dix}
			}
		}
		m[delta] = idx
	}
	return []int{-1, -1}
}
