package sprint



func SplitWhitespaces(s string) []string {

	var result []string
	var word string

	// if the string is empty, bye
	if s == "" {
		return nil
	}

	// for each value in the full string
	for _, r := range s {
		// Unless we get something weird
		if r > ' ' {
			// take whatever you get and put it in the string word
			word = word + string(r)
		// otherwise, we got something weird that needs to stop the word
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


