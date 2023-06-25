func longestPalindrome(s string) int {

	arr := make([]int, 52)
	ans := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			arr[s[i]-65]++
			if arr[s[i]-65] == 2 {
				ans += 2
				arr[s[i]-65] = 0
			}
		} else {
			arr[s[i]-71]++
			if arr[s[i]-71] == 2 {
				ans += 2
				arr[s[i]-71] = 0
			}
		}
	}

	for i := 0; i < 52; i++ {
		if arr[i] == 1 {
			ans++
			break
		}
	}

	return ans
}