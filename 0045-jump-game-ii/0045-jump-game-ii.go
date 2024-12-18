func jump(nums []int) int {

	// for i := len(nums) - 1; i <= 0; i-- {
	// 	if i+nums[i] >= len(nums)-1 {
	// 		b = true
	// 	}
	// }

	r := 0
	b := false
	min := 0
	_ = r
	_ = b
	_ = min

	arr := make([]int, len(nums))
	fmt.Println(arr)

	for i := 0; i < len(nums); i++ {

	forLoop:
		for j := 1; j <= nums[i]; j++ {
			if i+j > len(nums)-1 {
				break forLoop
			}

			if arr[i+j] == 0 {
				arr[i+j] = arr[i] + 1
			} else {
				if arr[i+j] > arr[i]+1 {
					arr[i+j] = arr[i] + 1
				}
			}
		}

		// r = i + nums[i]
		// if i+nums[i] >= len(nums)-1 {
		// 	b = true
		// }

	}

	// if arr[len(nums)-1]==0{
	// 	return -1
	// }
	return arr[len(nums)-1]

}