//go:build ignore
// +build ignore

package main

func NRune(s string, i int) rune {
	var char []rune
	for _, r := range s {
		char = append(char, r)
	}
	return char[i]
}