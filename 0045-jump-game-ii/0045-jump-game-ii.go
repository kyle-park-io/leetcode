func jump(nums []int) int {

	// b := false
	// for i := len(nums) - 1; i <= 0; i-- {
	// 	if i+nums[i] >= len(nums)-1 {
	// 		b = true
	// 	}
	// }

	arr := make([]int, len(nums))

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

			if i+j == len(nums)-1 {
				return arr[len(nums)-1]
			}
		}

	}

	return arr[len(nums)-1]
}