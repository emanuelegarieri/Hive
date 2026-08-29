//go:build ignore
// +build ignore

package main

func ShiftBy(r rune, step int) rune {
	position := int(r - 'a')
	new_position := int(position + step) % 26	
	return rune(new_position + 'a')
}
