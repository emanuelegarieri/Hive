//go:build ignore
// +build ignore

package main

func StrReverse(s string) string {

	var list []rune
	result := ""

	for _, r := range s {
		list = append(list, r)
	}

	for r := 1; r <= len(list); r++ {
		result = result + string(list[len(list)-r])
	}

	return result
}