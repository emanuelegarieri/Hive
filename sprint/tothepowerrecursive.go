//go:build ignore
// +build ignore

package main

func ToThePowerRecursive(n int, power int) int {

	switch power {
	case 0:
		return 1
	case 1:
		return n
	default:
		if power == 2 {
			return n * n
		} else {
			if power%2 == 0 {
				return ToThePowerRecursive(ToThePowerRecursive(n, power/2), 2)
			} else {
				return n * (ToThePowerRecursive(n, power-1))
			}
		}
	}
}