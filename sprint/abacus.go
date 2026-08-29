//go:build ignore
// +build ignore

package main

import "fmt"

func Abacus(a int, b int) int { 
	return a / b 
}

func main() {
	fmt.Println(Abacus(10, 5))
}

