package reverseinteger

import "math"

func reverse(x int) int {
	var x32 int32 = int32(x)

	if x32 == 0 {
		return 0
	}

	negative := x32 < 0

	if negative {
		x32 *= -1
	}

	for x32%10 == 0 {
		x32 /= 10
	}

	var res int32 = 0
	for x32 > 0 {
		lastDigit := x32 % 10
		if math.MaxInt32-res > lastDigit {
			res += lastDigit
		} else {
			return 0
		}
		if x32 >= 10 {
			if res <= math.MaxInt32/10 {
				res *= 10
			} else {
				return 0
			}
		}

		x32 /= 10
	}

	if negative {
		res *= -1
	}

	return int(res)
}
