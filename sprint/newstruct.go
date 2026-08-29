//go:build ignore
// +build ignore

package main

// Define the structure and its fields
type Point struct {
	X  		float32
	Y 		float32
	Text   	string
}

func MakePoint(x, y float32, text string) Point {
	// return the Point with assigned valiables to the fields
	return Point{
		X:  	x,
		Y: 		y,
		Text:   text,
	}
}