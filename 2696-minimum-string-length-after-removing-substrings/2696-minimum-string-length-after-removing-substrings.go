func minLength(s string) int {
	for {
		check := 0
		for i := 0; i < len(s)-1; i++ {
			var a int64
			b := int64(s[i+1])
			c := int64(s[i])
			a = int64(b - c)
			if a == 1 {
				if s[i] == 65 || s[i] == 67 {
					s = s[:i] + s[i+2:]
					check = 1
					break
				}
			}
		}
		if check == 0 {
			break
		} else {
			check = 0
		}
	}
	return len(s)
}
