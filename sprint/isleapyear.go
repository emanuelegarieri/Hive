//go:build ignore
// +build ignore

package main


// A leap year is a year that is evenly divisible by 4, except for years that are divisible by 100 but not by 400.

func IsLeapYear(year int) bool {
	if year%400 != 0 && year%100 == 0 {
		return false
	}
	return year%4 == 0
}