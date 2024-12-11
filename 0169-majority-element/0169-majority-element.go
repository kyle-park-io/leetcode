func majorityElement(nums []int) int {
	ans := 0
	count := 0
	arr := make(map[int]int)

	for _, num := range nums {
		arr[num]++
		if arr[num] > count {
			count = arr[num]
			ans = num
		}
	}

	fmt.Println(ans)
	return ans
}
