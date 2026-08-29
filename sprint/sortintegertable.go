//go:build ignore
// +build ignore

package main


func SortIntegerTable(table []int) []int {

	var result []int

	// for each number in the table
	for i := range table {
		// for each other number
		for j := i + 1; j < len(table); j++ {
			// if the previous number is bigger than the next number
			if table[i] > table[j] {
				// swap them
				table[j], table[i] = table[i], table[j]
			}
		}
		// append then the smaller
		result = append(result, table[i])
	}
	return result
}