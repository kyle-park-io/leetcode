func longestCommonPrefix(strs []string) string {

	if len(strs) == 1 {
		return strs[0]
	}

	if strs[0] == "" {
		return ""
	}
	criteron := strs[0]

	ans := ""
	for i := 0; i < len(criteron); i++ {
		c := true
		for j := 1; j < len(strs); j++ {
			if strs[j] == "" {
				return ""
			}

			if !c {
				break
			}

			if len(strs[j]) < i+1 {
				return ans
			}

			for k := 0; k <= i; k++ {
				if criteron[k] != strs[j][k] {
					c = false
					break
				}
			}

		}

		if c {
			ans = criteron[:i+1]
		}
	}

	return ans
}