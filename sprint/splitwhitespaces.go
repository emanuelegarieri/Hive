//go:build ignore
// +build ignore

package main



func SplitWhitespaces(s string) []string {

	var result []string
	var word string

	// if the string is empty, bye
	if s == "" {
		return nil
	}

	// for each value in the full string
	for _, r := range s {
		// If the character is not a whitespace
		if r > ' ' {
			// Add the character to the current word
			word = word + string(r)
		// Otherwise, end the current word
		} else {
			// if the word is not empty, append it to the final result  
			if word != "" {
				result = append(result, word)
				// reset the word variable
				word = ""
			}
		}
	}
	// if I have only one word, append it anyway
	if word != "" {
		result = append(result, word)
	}
	return result
}


