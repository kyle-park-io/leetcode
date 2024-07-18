func maxDepth(s string) int {
	l_max := 0
	plus := 0
	minus := 0
	result := 0

	for _, val := range s {
		if val == '(' {
			l_max++
		}
	}

	for _, val := range s {
		if val == '(' {
			plus++
		}
		if val == ')' {
			minus++
		}

		if l_max == plus {
			result = plus - minus
			break
		}
	}
	plus = 0
	minus = 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			plus++
		}
		if s[i] == '(' {
			minus++
		}
		if l_max == plus {
			if result < plus-minus {
				result = plus - minus
			}
			break
		}
	}

	return result
}

// // wrong answer
// func maxDepth(s string) int {
// 	l := 0
// 	l_ok := false
// 	r := 0
// 	r_ok := false

// 	for i := 0; i < len(s); i++ {

// 		if s[i] == '(' && !l_ok {
// 			l++
// 		}
// 		if s[i] == ')' {
// 			l_ok = true
// 		}

// 		if s[len(s)-i-1] == ')' && !r_ok {
// 			r++
// 		}
// 		if s[len(s)-i-1] == '(' {
// 			r_ok = true
// 		}

// 	}

// 	if l >= r {
// 		return l
// 	} else {
// 		return r
// 	}
// }

// // wrong answer
// func maxDepth(s string) int {
// 	l_max := 0
// 	l := 0
// 	r := 0

// 	for _, val := range s {
// 		if val == '(' {
// 			l_max++
// 		}
// 	}

// 	for _, val := range s {
// 		if val == '(' {
// 			l++
// 		}
// 		if val == ')' {
// 			r++
// 		}
// 		if l_max == l {
// 			break
// 		}
// 	}

// 	return l_max - r
// }