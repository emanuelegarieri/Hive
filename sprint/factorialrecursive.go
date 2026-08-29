//go:build ignore
// +build ignore

package main

func FactorialRecursive(n int) int {

	if n <= 1 {
		return 1
	}
	// 5 * 4 * 3 * 2  
	return n*FactorialRecursive(n-1) 
}	 