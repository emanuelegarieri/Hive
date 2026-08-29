//go:build ignore
// +build ignore

package main


// Create a function that takes a positive integer as input and returns the corresponding number of letters from the Latin alphabet. 
// The input integer won't be larger than the alphabet's length.
func AlphabetMastery(n int) string {

	// Create empty string
	text := ""

	// for each letter in rune append to the string "text" the corrisponding letter as a string starting from rune 'a'
	for i := 0; i < n; i++ {
		text += string(rune(letter + i))
	}
	// When the for loop is over, the funcion returns the "text" as a string
	return text
}



