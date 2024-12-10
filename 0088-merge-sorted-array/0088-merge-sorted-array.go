func merge(nums1 []int, m int, nums2 []int, n int) {

	result := make([]int, m+n)
	// if n == 0 {
	// 	result = nums1[:m]
	// 	// nums1 = result
	// 	copy(nums1, result)
	// 	return
	// }

	trial := 0
	trial_x := 0
	trial_y := 0
	for {
		if trial == m+n {
			// nums1 = result
			copy(nums1, result)
			fmt.Println(nums1)
			return
		}

		if trial_y >= n {
			if trial_x < m {
				result[trial] = nums1[trial_x]
				trial_x++
			}

			trial++
			continue
		}

		if trial_x >= m {
			result[trial] = nums2[trial_y]
			trial_y++
		} else {
			if nums1[trial_x] < nums2[trial_y] {
				result[trial] = nums1[trial_x]
				trial_x++
			} else {
				result[trial] = nums2[trial_y]
				trial_y++
			}
		}

		trial++
	}

}