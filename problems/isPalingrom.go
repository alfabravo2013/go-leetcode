package problems

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	digits := make([]int, 0)

	number := x
	for number > 0 {
		rem := number % 10
		digits = append(digits, rem)
		number = number / 10
	}

	var l, r = 0, len(digits) - 1
	for l < r {
		if digits[l] != digits[r] {
			return false
		}
		l++
		r--
	}

	return true
}
