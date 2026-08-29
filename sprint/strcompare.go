//go:build ignore
// +build ignore

package main

// 0 if the strings are equal,
// -1 if a comes before b in lexicographic order,
// 1 if a comes after b in lexicographic order.

func StrCompare(a, b string) int {

	rFirst := []rune(a)
	rSecond := []rune(b)

	// Handle edge cases
	if a == b {
		return 0
	}
	if a == "" {
		return -1
	}
	if b == "" {
		return 1
	}

	// To simplify the comparison, treat the longer string as the reference if needed
	if len(b) > len(a) {
		a, b = b, a
	}

	// For each letter of the shorter (b)
	for r := range len(rSecond) {
		// if the letter in position r of the a comes before, then give me -1
		if rFirst[r] < rSecond[r] {
			return -1
		// if it's comes after, then give me 1
		} else if rFirst[r] > rSecond[r] {
			return 1
		}
	}
	// Otherwise, we know that the word a continues and so b comes before a
	return 1
}
