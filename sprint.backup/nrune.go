package sprint

func NRune(s string, i int) rune {
	var char []rune
	for _, r := range s {
		char = append(char, r)
	}
	return char[i]
}