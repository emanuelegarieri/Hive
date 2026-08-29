//go:build ignore
// +build ignore

package main

// Write a Go function that takes an array of booleans and adds the minimum number of booleans necessary so that the count of true and false values in the array is equal. 
// The function should return the resulting modified array. 
// The order of the elements must remain the same and new elements should be added at the end of the array.

func BalanceOut(arr []bool) []bool {

	var result []bool
	var trues int
	var falses int
	var difference int
	
	if len(arr) == 0 {
		return nil
	}

	for i := range arr {
		if arr[i] == true {
			trues++
			result = append(result, arr[i])
		} else {
			falses++
			result = append(result, arr[i])
		}
	}

	difference = trues - falses

	if difference < 0 {
		for range -difference {
			result = append(result, true)
		}
	} else if difference > 0 {
		for range difference {
			result = append(result, false)
		}
	} else {
		return result
	}
	return result
}


// BalanceOut([]bool{true, false, false, false})
// >> [true false false false true true]