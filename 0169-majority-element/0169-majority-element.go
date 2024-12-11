func majorityElement(nums []int) int {
    var ans int
    var count int

    for _, num := range nums {
        if count == 0 {
            ans = num
        }
        if num == ans {
            count++
        } else {
            count--
        }
    }

    return ans

    // use map
	// ans := 0
	// count := 0
	// arr := make(map[int]int)

	// for _, num := range nums {
	// 	arr[num]++
	// 	if arr[num] > count {
	// 		count = arr[num]
	// 		ans = num
	// 	}
	// }

	// fmt.Println(ans)
	// return ans
}
