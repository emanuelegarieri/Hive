package sprint

import "fmt"

// Create a Go function that generates a series of unique triplets of digits. 

func Combinations() string {

	result := ""

	// for the combination of a ending to 7 (because the max number must be 789), b and c
	for a := 0; a <= 7; a++ {
		for b:= a + 1; b <= 8; b++ {
			for c := b+1; c <= 9; c++ {
				// if the loop is still going, so it's not empty, add the ", " at the end of the triplet 
				// The first round of loop is empty, so it skips the concatenation of ", "
				if result != "" {
					result += ", "
				}
				// Add the triplet as a string, it must be inside the for loop, otherwise ", " are added also to the last triplet.
				// ", " are not added when the for loop breaks
				result += fmt.Sprintf("%1d%1d%1d", a, b, c)
			}
		}
	}
	return result
}