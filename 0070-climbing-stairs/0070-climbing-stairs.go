func repeat(i int, n int, memo []int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	if memo[i] > 0 {
		return memo[i]
	}
	memo[i] = repeat(i+1, n, memo) + repeat(i+2, n, memo)
	return memo[i]
}

func climbStairs(n int) int {
	memo := make([]int, n+1)
	ans := repeat(0, n, memo)
	return ans
}