//go:build ignore
// +build ignore

package main

func RuneChecksum(a rune, b rune) rune{
	new_pos_a := int(a -'a') +1
	new_pos_b := int(b -'a') +1
	multip := int(new_pos_a * new_pos_b)
	res := (multip-1) % 26 +1
	rune_res := rune((res -1) + int('a'))
	return rune_res
}