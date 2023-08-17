package fibonacci

func Fibonacci(n int) []int {
	var results []int

	for i := 0; i < n; i++ {
		if i <= 1 {
			results = append(results, i)
		} else {
			nextNum := results[i-1] + results[i-2]
			results = append(results, nextNum)
		}

	}
	return results
}
