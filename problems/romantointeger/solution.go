package romantointeger

func romanToInt(s string) int {
	symbols := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	number := 0
	curDigit := 0
	prevDigit := 0
	for i, r := range s {
		curDigit = symbols[r]
		if i > 0 && curDigit > prevDigit {
			number -= prevDigit
			number += curDigit - prevDigit
		} else {
			number += curDigit
		}
		prevDigit = curDigit
	}

	return number
}
