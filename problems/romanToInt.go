package problems

var romanNumbers = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	sum, prev := 0, 0
	runes := []rune(s)
	for i, r := range runes {
		value, ok := romanNumbers[r]
		if !ok {
			return -1
		}

		sum += value

		if i > 0 {
			prev = romanNumbers[runes[i-1]]
		}

		if prev > 0 && prev < value {
			sum -= prev * 2
		}
	}
	return sum
}
