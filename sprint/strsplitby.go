package sprint


func StrSplitBy(s, sep string) []string {

	var result []string
	var word string
	var check string

	rString := []rune(s)
	rSep := []rune(sep)

	// if the string is empty, bye
	if s == "" {
		return nil
	}
	if sep == "" {
		return []string{s}
	}

	// First we change the separator in values we can recognise (IMPORTANT, WE DEFINE THE JUMPS IN THE FOR LOOP IN THIS CASE)
	for r := 0; r < len(rString); {
		// check restart to be empty
		check = ""
		if rString[r] == rSep[0] {
			// for the lenght of the separator
			for i := 0; i < len(rSep); i++ {
				// until it's valid, we create a word long as much as the separator
				if r+i < len(rString) {
					check = check + string(rString[r+i])
				}
			}
		}
		// then we check if the word and the separator are equal
		if check == sep {
			// if we have had a word
			if word != "" {
				// append it
				result = append(result, word)
				// and reset it
				word = ""
			}
			// (THEN WE MOVE R TO THE NEXT REAL POSITION)
			r = r + len(rSep)
		// if the check don't pass, so we were mistaken	
		} else {
			// we want the word
			word = word + string(rString[r])
			// (OTHER STEP IS DEFINED IN HERE AND WE WANT IT BY ONE) and we move ahead
			r++
		}
	}
	// if I have only one word, append it anyway
	if word != "" {
		result = append(result, word)
	}
	return result
}

