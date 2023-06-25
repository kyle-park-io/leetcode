func canJump(nums []int) bool {

	check_i := 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		if i > check_i {
			ans = 1
			break
		}

		if i+nums[i] > check_i {
			check_i = i + nums[i]
		}
	}

	if ans == 1 {
		return false
	} else {
		return true
	}
}