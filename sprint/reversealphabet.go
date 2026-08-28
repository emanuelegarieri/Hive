package sprint

// Create a function that takes a step value as input. 
// Starting from z in the Latin alphabet, it should return the lowercase alphabet in reverse order as a string. 
// For each letter, you skip step-1 letters. If step is 0 or negative, use a default value of 1.

func ReverseAlphabet(step int) string {

	// Add the condition to the parameter step, that if it's 0 or negative, then it's set to 1
	if step <=0 {
		step = 1
	}
	
	// Create an empty variable string called "reverse"
	reverse := ""

	// Start a for loop starting from rune 'z' to rune 'a', converted as integers for allign them to the integer parameter step, and moving backward with step based on the parameter step
	for r := int('z'); r >= int('a'); r -= step {
		reverse += string(rune(r))
	}

	// Return the string called "reverse"
	return reverse
}