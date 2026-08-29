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
	return Point {
		X:  	x,
		Y: 		y,
		Text:   text,
	}
}


func PointDiff(p1, p2 Point) Point {
	// This function gets 2 Points created using the func MakePoint and called p1 and p2
	// Then we check the fields of p1 and p2 and we do our things
	if (p1.X + p1.Y) >= (p2.X + p2.Y) {
		return p1
	} else {  
		return p2
	} 
}