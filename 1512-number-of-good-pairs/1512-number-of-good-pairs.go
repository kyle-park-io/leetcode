func numIdenticalPairs(nums []int) int {
	map_n := make(map[int]int)
	result := 0

	for _, val := range nums {
		map_n[val]++
	}

	// combination
	for _, val := range map_n {
		result += (val) * (val - 1) / 2

	}

	return result
}