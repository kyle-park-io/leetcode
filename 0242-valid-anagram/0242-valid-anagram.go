func isAnagram(s string, t string) bool {
	// UTF-8(ASCII)
	// a~z (97~122)
	s_arr1 := [26]int{}
	t_arr1 := [26]int{}
	// arr2 := make([]int, 26)
	// arr3 := make(map[int]struct{})

	for _, v := range s {
		s_arr1[v-97]++
	}
	for _, v := range t {
		t_arr1[v-97]++
	}
	for i := 0; i < len(s_arr1); i++ {
		if s_arr1[i] != t_arr1[i] {
			return false
		}
	}

	return true
}