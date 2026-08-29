//go:build ignore
// +build ignore

package main


// Write a Go function that takes an array of float64, two indices, and removes elements that lie between these indices from the array. 
// The lower index should be removed, and the upper index should be kept. The function should return the resulting modified array. 
// The indices can be negative or larger than the length of the array, but the function should still remove the elements that fall within the specified range. 
// The indices might also be in wrong order.


func RemoveElementsInRange(arr []float64, from, to int) []float64 {

	// Except for from bigger than to
	if from > to {
		from, to = to, from
	}
	// Except for to lager than to
	if to > len(arr) {
		to = len(arr)
	}
	// Except for from less than 0
	if from < 0 {
		from = 0
	}
	// If the array length is 0, return nil
	if len(arr) == 0 {
		return nil
	}
	var slice []float64

	// for each float
	for i := 0; i < len(arr); i++ {
		// If the corresponding position of arr is not in the range between from and to
		if i < from || i >= to {
			// slice gets the float in the position position
			slice = append(slice, arr[i])
		}
	}
	return slice
}

// RemoveElementsInRange([]float64{10, 9, 8, 7, 6, 5}, 1, 4)
// >> [10 6 5]