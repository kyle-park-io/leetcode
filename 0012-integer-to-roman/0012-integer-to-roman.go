func intToRoman(num int) string {
	arr1 := []string{"M", "D", "C", "L", "X", "V", "I"}
	arr2 := []uint16{1000, 500, 100, 50, 10, 5, 1}

	arr3 := [3][]string{{"IX", "V", "IV", "I"}, {"XC", "L", "XL", "X"}, {"CM", "D", "CD", "C"}}
	result := []string{}
	calc := uint16(num)
	a := int(uint16(num) / arr2[0])
	calc -= uint16(a) * arr2[0]
	for i := 0; i < a; i++ {
		result = append(result, arr1[0])
	}
	for i := 2; i >= 0; i-- {
		if calc >= 9*uint16(math.Pow10(i)) {
			result = append(result, arr3[i][0])
			calc -= 9 * uint16(math.Pow10(i))
		}
		if calc >= 5*uint16(math.Pow10(i)) {
			result = append(result, arr3[i][1])
			calc -= 5 * uint16(math.Pow10(i))
		}
		if calc >= 4*uint16(math.Pow10(i)) {
			result = append(result, arr3[i][2])
			calc -= 4 * uint16(math.Pow10(i))
		}
		if calc >= 1*uint16(math.Pow10(i)) {
			a := int(uint16(calc) / (1 * uint16(math.Pow10(i))))
			calc -= uint16(a) * 1 * uint16(math.Pow10(i))
			for j := 0; j < a; j++ {
				fmt.Println(arr3[i][3])
				result = append(result, arr3[i][3])
			}
		}
	}
	b := strings.Join(result, "")
	return b
}