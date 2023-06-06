func romanToInt(s string) int {
	map_roman := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0

	for i := 0; i < len(s)-1; i++ {
		str := string(s[i])
		str_next := string(s[i+1])
		check := false

		if str == "I" {
			if str_next == "V" {
				result += 4
				i++
				check = true
			} else if str_next == "X" {
				result += 9
				i++
				check = true
			}
		} else if str == "X" {
			if str_next == "L" {
				result += 40
				i++
				check = true
			} else if str_next == "C" {
				result += 90
				i++
				check = true
			}
		} else if str == "C" {
			if str_next == "D" {
				result += 400
				i++
				check = true
			} else if str_next == "M" {
				result += 900
				i++
				check = true
			}
		}
		if !check {
			result += map_roman[str]
			if i == len(s)-2 {
				result += map_roman[str_next]
			}
		} else {
			if i == len(s)-2 {
				result += map_roman[string(s[len(s)-1])]
			}
		}
		check = false
	}
	if len(s) == 1 {
		result += map_roman[string(s)]
	}
	return result
}