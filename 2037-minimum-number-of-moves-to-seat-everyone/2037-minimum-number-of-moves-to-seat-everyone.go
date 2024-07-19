func minMovesToSeat(seats []int, students []int) int {

	result := 0

	sort.Ints(seats)
	sort.Ints(students)

	for i := 0; i < len(seats); i++ {
		calc := seats[i] - students[i]
		result += int(math.Abs(float64(calc)))

	}

	return result
}
