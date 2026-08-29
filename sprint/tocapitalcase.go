//go:build ignore
// +build ignore

package main

func ToCapitalCase(s string) string {

	result := ""
	new_word := true

	for _, r := range s {
		// If r is a lowercase letter and the word is new, make it uppercase
		if r >= 'a' && r <= 'z' {
			if new_word == true {
				r = r + 'A' - 'a'
			}
			result += string(r)
			new_word = false
		// otherwise if the letter is uppercase and the word is not new, make it lowercase
		} else if r >= 'A' && r <= 'Z' {
			if new_word == false {
				r = r - 'A' + 'a'
			}
			result += string(r)
			new_word = false
		// if we get a number, append it, and the word is still old
		} else if r >= '0' && r <= '9' {
			result += string(r)
			new_word = false
		// but if it's a symbol the word is new
		} else {
			result += string(r)
			new_word = true
		}
	}
	return result
}