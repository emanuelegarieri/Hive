//go:build ignore
// +build ignore

package main

func SimpleStrToInt(s string) int {

	result := 0
	// Process each character in the string
	for n := 0; n < len(s); n++ {
		// if the value is included from 0 to 9, otherwise return 0
		if s[n] < '0' || s[n] > '9' {
			return 0
		}
		// the number as an integer of the value
		value := int(rune(s[n] - '0'))
		// becomes the temporary result, added to a starting empty sum. 
		// The previous temporary result is multiplied by 10 and we add the last found value.
		result = result*10 + value
	}

	return result
}

/* Proper way:
import (
	"fmt"
	"strconv"
)

func SimpleStrToInt(s string) int {

	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return result
} */
