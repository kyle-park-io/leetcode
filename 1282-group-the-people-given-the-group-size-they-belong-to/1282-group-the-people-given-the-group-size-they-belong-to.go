func groupThePeople(groupSizes []int) [][]int {
	member := make(map[int][]int)
	result := make([][]int, 0)

	for i, val := range groupSizes {
		member[val] = append(member[val], i)
	}

	for i, val := range member {

		q := int(math.Ceil(float64(len(val)) / float64(i)))
		for j := 1; j <= q; j++ {
			if j == q {
				subSlice := val[i*(j-1):]
				result = append(result, subSlice)
			} else {
				subSlice := val[i*(j-1) : i*j]
				result = append(result, subSlice)
			}
		}
	}

	return result
}