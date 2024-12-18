func canCompleteCircuit(gas []int, cost []int) int {

	n := len(gas)
	fuelLeft, globalFuelLeft, start := 0, 0, 0

	for i := 0; i < n; i++ {

		globalFuelLeft += gas[i] - cost[i]
		fuelLeft += gas[i] - cost[i]

		if fuelLeft < 0 {
			start = i + 1
			fuelLeft = 0
		}
	}

	if globalFuelLeft < 0 {
		return -1
	}

	return start
}

func canCompleteCircuit2(gas []int, cost []int) int {

	l := len(gas)
	g := 0
	c := 0

	for i := 0; i < len(gas); i++ {

	forLoop:
		for {
			if c == len(gas) {
				fmt.Println(i)
				return i
			}

			g += gas[(i+c)%l]

			if g < cost[(i+c)%l] {
				break forLoop
			}

			g -= cost[(i+c)%l]
			c++
		}

		g = 0
		c = 0

	}

	return -1
}
