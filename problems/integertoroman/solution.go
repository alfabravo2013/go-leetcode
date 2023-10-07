package integertoroman

func intToRoman(num int) string {
	res := []rune{}

	for num >= 1000 {
		res = append(res, 'M')
		num -= 1000
	}

	if num >= 900 {
		res = append(res, 'C')
		res = append(res, 'M')
		num -= 900
	}

	if num >= 500 {
		res = append(res, 'D')
		num -= 500
	}

	if num >= 400 {
		res = append(res, 'C')
		res = append(res, 'D')
		num -= 400
	}

	for num >= 100 {
		res = append(res, 'C')
		num -= 100
	}

	if num >= 90 {
		res = append(res, 'X')
		res = append(res, 'C')
		num -= 90
	}

	if num >= 50 {
		res = append(res, 'L')
		num -= 50
	}

	if num >= 40 {
		res = append(res, 'X')
		res = append(res, 'L')
		num -= 40
	}

	for num >= 10 {
		res = append(res, 'X')
		num -= 10
	}

	if num >= 9 {
		res = append(res, 'I')
		res = append(res, 'X')
		num -= 9
	}

	if num >= 5 {
		res = append(res, 'V')
		num -= 5
	}

	if num >= 4 {
		res = append(res, 'I')
		res = append(res, 'V')
		num -= 4
	}

	for num > 0 {
		res = append(res, 'I')
		num -= 1
	}

	return string(res)
}
