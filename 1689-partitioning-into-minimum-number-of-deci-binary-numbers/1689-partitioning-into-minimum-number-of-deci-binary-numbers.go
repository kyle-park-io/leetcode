func minPartitions(n string) int {

	// 0: 48 9: 57
	b := []byte(n)
	for i := 0; i < len(b); i++ {
		b[i] -= 48
	}

	check := 0
	num := 0

	for {
		for i := 0; i < len(b); i++ {
			if b[i] > 0 {
				b[i] -= 1
			} else {
				check++
			}
		}

		if check == len(b) {
			break
		}

		check = 0
		num++
	}

	// slice sort
	// sort.Slice(b, func(i, j int) bool {
	// 	return b[i] < b[j]
	// })
	// fmt.Println(b)

	// for i, val := range b {
	// 	result += int(val) - 48
	// 	for j := i; j < len(n); j++ {
	// 		if b[j] <= b[i] {
	// 			b[j] -= b[i]
	// 		}
	// 	}
	// }

	return num
}