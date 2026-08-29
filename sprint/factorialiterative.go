//go:build ignore
// +build ignore

package main

func FactorialIterative(n int) int {

	result := 1
	for i := 1; i <= n; i++ {
		result = i * result
	}
	return result
}