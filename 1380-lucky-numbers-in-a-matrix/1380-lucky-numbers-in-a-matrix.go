func luckyNumbers(matrix [][]int) []int {

	check := 0
	max := int(math.Pow(10, 5)) + 1
	x := -1
	check_x := make([]bool, len(matrix))
	result := make([]int, 0)

	for i := 0; i < len(matrix[0]); i++ {

		for j := 0; j < len(matrix); j++ {
			if check < matrix[j][i] {
				x = j
				check = matrix[j][i]
			}
		}

		for j := 0; j < len(matrix[0]); j++ {
			if max > matrix[x][j] {
				max = matrix[x][j]
			}
		}

		if check == max && !check_x[x] {
			check_x[x] = true
			result = append(result, check)
		}

		check = 0
		max = int(math.Pow(10, 5)) + 1
		x = -1

	}

	return result
}

// func luckyNumbers(matrix [][]int) []int {

// 	check := 0
// 	max := int(math.Pow(10, 5)) + 1
// 	x := -1
// 	result := make([]int, 0)

// 	for i := 0; i < len(matrix[0]); i++ {

// 		for j := 0; j < len(matrix); j++ {
// 			if check < matrix[j][i] {
// 				x = j
// 				check = matrix[j][i]
// 			}
// 		}

// 		for j := 0; j < len(matrix[0]); j++ {
// 			if max > matrix[x][j] {
// 				// y = j
// 				max = matrix[x][j]
// 			}
// 		}

// 		if check == max {
// 			result = append(result, check)
// 		}

// 		check = 0
// 		max = int(math.Pow(10, 5)) + 1
// 		x = -1

// 	}

// 	return result
// }