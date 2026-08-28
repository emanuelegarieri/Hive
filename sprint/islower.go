package sprint

func IsLower(s string) bool {

	var list []rune

	if len(s) == 0 {
		return false
	}
	
	for _, ch := range s {
		list = append(list, ch)
	}

	for _, r := range list {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}