package string

func ReverseAlphabetValue(ch rune) rune {
	position := int(ch - 'a')
	new_position := rune(25 - position) +'a'
	return rune(new_position)

}