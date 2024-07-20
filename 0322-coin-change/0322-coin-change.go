func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)

	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	for i := 0; i < len(dp); i++ {
		for _, val := range coins {
			if i+val > amount {
				continue
			}

			if dp[i+val] > dp[i]+1 {
				dp[i+val] = dp[i] + 1
			}
		}
	}
	fmt.Println(dp)

	if dp[len(dp)-1] == amount+1 {
		return -1
	}
	return dp[len(dp)-1]
}

// // wrong answer
// func coinChange(coins []int, amount int) int {

// 	result := -1
// 	current := 0

// 	dp := make([]int, amount+1)
// 	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

// 	if amount == 0 {
// 		return 0
// 	}

// 	for _, val := range coins {
// 		for {
// 			if current+val > amount {
// 				break
// 			}
// 			if dp[current+val] > dp[current]+1 || dp[current+val] == 0 {
// 				dp[current+val] = dp[current] + 1
// 			}
// 			current += val
// 		}
// 		current = 0
// 	}
// 	fmt.Println(dp)

// 	for i := 0; i < len(dp); i++ {
// 		if result == -1 {
// 			result = (dp[i] + dp[amount-i])
// 		} else if dp[i] != 0 {
// 			result = int(math.Min(float64(result), float64(dp[i]+dp[amount-i])))
// 		}
// 	}

// 	if dp[amount] == 0 {
// 		return -1
// 	}

// 	return result
// }