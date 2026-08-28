package sprint

func GetLastRune(s string) rune {
	var last []rune
	for _, c := range s {
		last = append(last, c)
	}
	return last[len(last)-1]
}
