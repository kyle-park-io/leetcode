func countStudents(students []int, sandwiches []int) int {

	i := 1
	result := -1
	check := 0

	for {
		fmt.Printf("index : %d\n", i)

		if len(students) == 0 {
			result = 0
			break
		} else {
			for _, val := range students {
				if sandwiches[0] == val {
					check++
				}
			}
			if check == 0 {
				result = len(students)
				break
			}
		}

		b := pick(students, sandwiches)
		if b {
			students = students[1:]
			sandwiches = sandwiches[1:]

			fmt.Printf("students: %v\n", students)
			fmt.Printf("sandwiches: %v\n", sandwiches)
			// time.Sleep(time.Second * 3)
			continue
		}

		students = sort(students)

		fmt.Printf("students: %v\n", students)
		fmt.Printf("sandwiches: %v\n", sandwiches)
		// time.Sleep(time.Second * 3)

		check = 0
		i++
	}

	return result
}

func sort(students []int) []int {
	students = append(students[1:], students[0])
	return students
}

func pick(students []int, sandwiches []int) bool {
	if students[0] == sandwiches[0] {
		return true
	}
	return false
}