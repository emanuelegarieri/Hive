//go:build ignore
// +build ignore

package main

// Write a function that takes an integer n and returns true if n is a prime number, and false otherwise. 
// A prime number is a natural number greater than 1 that has no divisors other than 1 and itself.

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}