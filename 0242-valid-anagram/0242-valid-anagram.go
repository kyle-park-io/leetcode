func isAnagram(s string, t string) bool {

	// UTF-8(ASCII)
	// a~z (97~122)
	// s
	s_arr1 := [26]int{}
	s_arr2 := make([]int, 26)
	s_arr3 := make(map[rune]int)
	_ = s_arr1
	_ = s_arr2
	_ = s_arr3
	// t
	t_arr1 := [26]int{}
	t_arr2 := make([]int, 26)
	t_arr3 := make(map[rune]int)
	_ = t_arr1
	_ = t_arr2
	_ = t_arr3
	// map (a~z)
	map_all := make(map[rune]struct{})
	_ = map_all

	// // case1:
	// for _, v := range s {
	// 	s_arr1[v-97]++
	// 	// s_arr2[v-97]++
	// }
	// for _, v := range t {
	// 	t_arr1[v-97]++
	// 	// t_arr2[v-97]++
	// }
	// for i := 0; i < len(s_arr1); i++ {
	// 	if s_arr1[i] != t_arr1[i] {
	// 		return false
	// 	}
	// 	// if s_arr2[i] != t_arr2[i] {
	// 	// 	return false
	// 	// }
	// }
	// return true

	// // case2:
	// for _, v := range s {
	// 	s_arr3[v-97]++
	// 	map_all[v-97] = struct{}{}
	// }
	// for _, v := range t {
	// 	t_arr3[v-97]++
	// 	map_all[v-97] = struct{}{}
	// }
	// for i := range map_all {
	// 	if s_arr3[i] != t_arr3[i] {
	// 		return false
	// 	}
	// }
	// return true

	// case3:
	if len(s) != len(t) {
		return false
	}
	copied_s := []rune(s)
	copied_t := []rune(t)
	sort.Slice(copied_s, func(i, j int) bool {
		return copied_s[i] < copied_s[j]
	})
	sort.Slice(copied_t, func(i, j int) bool {
		return copied_t[i] < copied_t[j]
	})
	for i := 0; i < len(copied_s); i++ {
		if copied_s[i] != copied_t[i] {
			return false
		}
	}
	return true
}