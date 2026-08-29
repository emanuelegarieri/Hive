//go:build ignore
// +build ignore

package main


func StrToInt(s string) int {

	result := 0
	runes := []rune(s)
	var sign bool

	for i, n := range runes {
		if i == 0  && (n == '+' || n == '-') {
			if n == '-' {
				sign = true
			}
			continue
		} 
		if n < '0' || n > '9' {
			return 0
		}	
		result = result*10 + int(n-'0')
	}

	if sign {
		result = -result
	}
	return result
}