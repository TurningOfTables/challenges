package matchinttotal

// MatchIntTotal takes an input slice of ints, and returns the total of any ints which appear more than once
func MatchIntTotal(input []int) int {
	inputData := make(map[int]int)
	var inputTotal int

	// Gather data on how many times each int appears
	for _, v := range input {
		inputData[v]++
	}

	// If an int appeared more than once, add (int * frequency) to the total
	for k, v := range inputData {
		if v > 1 {
			inputTotal += (v * k)
		}
	}
	return inputTotal
}
