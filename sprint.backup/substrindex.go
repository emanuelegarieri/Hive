package sprint


func SubstrIndex(s string, toFind string) int {

	var check string
	var count int

	rString := []rune(s)
	rSep := []rune(toFind)

	// if the string or the index to find is empty, bye
	if toFind == "" {
		return 0
	}
	if s == "" {
		return -1
	}

	// First we change the separator in values we can recognise (IMPORTANT, WE DEFINE THE JUMPS IN THE FOR LOOP IN THIS CASE)
	for r := range len(rString) {
		// check restart to be empty
		check = ""
		count++
		if rString[r] == rSep[0] {
			// for the lenght of the separator
			for i := range len(rSep) {
				// until it's valid, we create a word long as much as the separator
				if r+i < len(rString) {
					check = check + string(rString[r+i])
				}
			}
		}
		// then we check if the word and the separator are equal
		if check == toFind {
			return count - 1
		} else {
			check = ""
		}
	}
	return -1
}