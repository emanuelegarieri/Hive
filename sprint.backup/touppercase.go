package sprint

func ToUpperCase(s string) string {
	var list []rune
	var result string

	if s == "" {
		return s
	}
	
	for _, ch := range s {
		list = append(list, ch)
	}

	for _, r := range list {
		if r < 'a' || r > 'z' {
			result = result + string(r)
		} else {
			result = result + string(r+'A'-'a')
		}
	}
	return result
}
