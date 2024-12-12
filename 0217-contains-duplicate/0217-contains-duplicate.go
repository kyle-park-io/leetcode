func containsDuplicate(nums []int) bool {

	arr := make(map[int]bool)

	for _, v := range nums {
		if arr[v] {
			return true
		}
		arr[v] = true
	}
	return false
}
