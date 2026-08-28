package game

import "strings"


func CheckGuess(word string, guess string, letters string) (string, string) {

	green  	:= "\u001B[32m"
	yellow 	:= "\u001B[33m"
	white 	:= "\u001B[37m"
	reset  	:= "\u001B[0m"
	wordRunes := []rune(word)
	guessRunes := []rune(guess)

	// Create 2 lists with defaut values long 5
	result := make([]int, len(guessRunes))
	used := make([]bool, len(wordRunes))

	// First check the letters in the correct position
	for i := range guessRunes {
		// If match exactly
		if guessRunes[i] == wordRunes[i] {
			// Pop up the 0 to 2
			result[i] = 2
			// Mark it as used
			used[i] = true
		}
	}

	// If green, skip
	for i := range guessRunes {
		if result[i] == 2 {
			continue
		}
		// Then check the correct letters in the wrong position
		for j := range wordRunes {
			// Check each of the remaining false for the word letters
			if used[j] == false && guessRunes[i] == wordRunes[j] {
				// if match, yellow
				result[i] = 1
				used[j] = true
				// break because we already found it for that if, but it goes back to the loop for the next letter
				break
			}
		}
	}

	// Art attack
	feedback := ""
	for i := range guessRunes {
		// Feedback is always uppercase
		letter := strings.ToUpper(string(guessRunes[i]))
		if result[i] == 2 {
			feedback += green + letter + reset
		} else if result[i] == 1 {
			feedback += yellow + letter + reset
		} else {
			feedback += white + letter + reset
		}
	}

	letters = updateLetters(letters, guessRunes, result)
	return feedback, letters
}

func updateLetters(letters string, guessRunes []rune, result []int) string {
	var newLetters string
	for _, letter := range letters {
		var remove bool
		// lowcase it
		lowcase := rune(strings.ToLower(string(letter))[0])
		
		// for each letter in the guess word
		for i := range guessRunes {
			// if the letter was completely wrong
			if lowcase == guessRunes[i] && result[i] == 0 {
				found := false
				for j := range guessRunes {
					// If we guessed before
					if guessRunes[j] == lowcase && (result[j] == 1 || result[j] == 2) {
						// we found it and skip
						found = true
						break
					}
				}
				// if not found, we want to remove it
				if !found {
					remove = true
				}
			}
		}
		// then we append for those we do not want to remove
		if remove == false {
			newLetters += string(letter)
		}
	}
	return newLetters
}