func maxArea(height []int) int {

	left := 0
	right := len(height) - 1
	calc := 0
	ans := 0
	for {
		if left >= right {
			break
		}
		calc = int(math.Min(float64(height[left]), float64(height[right]))) * (right - left)
		if calc >= ans {
			ans = calc
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return ans
}