func findContentChildren(g []int, s []int) int {

	sort.Ints(g)
	sort.Ints(s)

	if len(s) == 0 {
		return 0
	}

	c := len(s) - 1
	result := 0
forLoop:
	for i := len(g) - 1; i >= 0; i-- {
		if c < 0 {
			break forLoop
		}

		if g[i] <= s[c] {
			result++
			c--
		}
	}

	fmt.Println(result)
	return result
}