//go:build ignore
// +build ignore

package main

func IsNumeric(s string) bool {

	var list []rune

	if len(s) == 0 {
		return false
	}
	
	for _, ch := range s {
		list = append(list, ch)
	}

	for _, r := range list {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true

}