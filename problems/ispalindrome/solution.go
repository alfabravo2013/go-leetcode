package ispalindrome

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	digits := getDigits(x)
	for i := 0; i < len(digits)/2; i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}

	return true
}

func getDigits(x int) []int {
	digits := make([]int, 0)
	for x > 0 {
		digit := x % 10
		digits = append(digits, digit)
		x = x / 10
	}
	return digits
}

func isPalindromeEfficient(x int) bool {
	if x < 0 {
		return false
	}

	acc := 0

	for n := x; n > 0; n /= 10 {
		lastDigit := n % 10
		acc *= 10
		acc += lastDigit
	}

	return x == acc
}
