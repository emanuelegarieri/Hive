//go:build ignore
// +build ignore

package main


// Create a function that takes a one-digit integer as input and returns a countdown string. 
// The countdown should start from the given number, skip every second number, and end with 0!. 
// For example, if the input is 7, the function should return "7, 5, 3, 1, 0!".

func Countdown(n int) string {

	countdown := ""
	
	for i := n; i > 0; i -= 2 {
		if countdown != "" {
			countdown += ", "
		}

		// We need to concat the int as a string, converting the correspondent rune from the integer starting from the rune '0' 
		// eg. if i == 5 as an integer, then we can find the correspondent value as as rune adding the starting rune '0'
		countdown += string(rune('0' + i))

	}

	if countdown != "" {
		countdown += ", "
	}
	countdown += "0!"

	return countdown
}

