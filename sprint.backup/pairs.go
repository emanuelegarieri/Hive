package sprint

import "fmt"

// Create a Go function that finds all possible combinations of two two-digit numbers. 


// Each number within a pair should be separated by a space. 
// The function should return a string containing these pairs.

func Pairs() string {

	result := ""

	for a := 0; a <= 98; a++ {
		for b:= a + 1; b <= 99; b++ {
			// Add a separator comma and a space for each pairs, skipping the last call. 
			if result != "" {
			result += ", "			
		} 

		// Here the pairs are added using Sprintf, so specifying how they need to be formatted.
		// In our case we use % 0(the 0 we need), 2 (to fill until having 2 digits), and 2 cause we want to limit our prints to 2 digits only.
		// Then, when the for loops end their job add the last 2 pairs, skipping then the comma and the space.
		result += fmt.Sprintf("%02d %02d", a, b)
	}
	}

	// The function returns the string combination of two two-digit numbers with a leading zero if it's less than 10.
	return result
}