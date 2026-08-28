package sprint

// The sum of elements is lower than the given value.

func FilterBySum(arr [][]int, limit int) [][]int {

	// I create a list of lists of integers
	var result [][]int

	// If the mother list is empty, ciao
	if len(arr) == 0 {
		return nil
	}

	// for each list in the mother list
	for i := range arr {
		sum := 0
		// If the list is empty, ciao
		if len(arr[i]) != 0 {
		// Otherwise I sum all the numbers of the list in the position is
		for j := range arr[i] {
			sum += arr[i][j]
		}
		// If the sum is bigger or equal than the limit
		if sum >= limit {
			// I append the list with index i to the result
			result = append(result, arr[i])
		}
		}
	}
	return result
}


