package main

// import "fmt"

type Rectangle struct {
	Width 		int
	Height 	 	int
	Area 	 	int
	Perimeter  	int
}

type Coords struct {
	X			int
	Y	 	 	int
}


func GetRectangle(min, max Coords) Rectangle {

	lenX := max.X-min.X
	lenY := max.Y-min.Y
	
	return Rectangle {
		Width: lenX,
		Height: lenY,
		Area: lenX*lenY,
		Perimeter: (lenX+lenY)*2,
	}
}

// func main() {
	
// 	min := Coords{X: 1, Y: 1}
// 	max := Coords{X: 6, Y: 4}
// 	r:= GetRectangle(min, max)
	
// 	fmt.Println(r.Width)
// 	fmt.Println(r.Height)
// 	fmt.Println(r.Area)
// 	fmt.Println(r.Perimeter)

// }

