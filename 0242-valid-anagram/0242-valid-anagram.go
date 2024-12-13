func isAnagram(s string, t string) bool {
	// UTF-8(ASCII)
	// a~z (97~122)

	// s
	s_arr1 := [26]int{}
	s_arr2 := make([]int, 26)
	s_arr3 := make(map[int]int)
	_ = s_arr1
	_ = s_arr2
	_ = s_arr3
	// t
	t_arr1 := [26]int{}
	t_arr2 := make([]int, 26)
	t_arr3 := make(map[int]int)
	_ = t_arr1
	_ = t_arr2
	_ = t_arr3

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