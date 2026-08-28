package main

import "fmt"

// Define the structure and its fields
type Point struct {
	X    float32
	Y    float32
	Text string
}


func MakePoint(x, y float32, text string) Point {
// return the Point with assigned valiables to the fields
	return Point{
		X:    x,
		Y:    y,
		Text: text,
	}
}


func PointText(p Point) Point {
	// We change the text of the given Point as a print adding its coordinates
	p.Text = fmt.Sprintf("Text at (%f, %f)", p.X, p.Y)
	return p
}