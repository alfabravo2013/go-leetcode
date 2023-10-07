package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	bytes := make([]byte, 0)

	if len(strs) == 0 {
		return string(bytes)
	}

	for idx := 0; idx < len(strs[0]); idx++ {
		r := strs[0][idx]
		for _, s := range strs {
			if idx > len(s)-1 || s[idx] != r {
				return string(bytes)
			}
		}
		bytes = append(bytes, r)
	}

	return string(bytes)
}
