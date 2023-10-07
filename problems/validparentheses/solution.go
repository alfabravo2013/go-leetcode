package validparentheses

func isValid(s string) bool {
	stck := make([]rune, 0)

	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			stck = append(stck, r)
		} else {
			if len(stck) == 0 {
				return false
			}
			if stck[len(stck)-1] != getOpening(r) {
				return false
			}
			stck = stck[:len(stck)-1]
		}
	}

	return len(stck) == 0
}

func getOpening(r rune) rune {
	switch r {
	case '}':
		return '{'
	case ']':
		return '['
	case ')':
		return '('
	default:
		return ' '
	}
}
