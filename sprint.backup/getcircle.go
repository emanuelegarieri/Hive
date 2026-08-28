package main

// import "fmt"

type Circle struct {
	Radius 		float32
	Diameter  	float32
	Area  		float32
	Perimeter  	float32
}


func GetCircle(r float32) Circle {

	return Circle {
		Radius: r,
		Diameter: r*2,
		Area: r*r*3.14,
		Perimeter: r*2*3.14,
	}
}

// func main() {
	
// 	c:= GetCircle(5)
// 	fmt.Println(c.Radius)
// 	fmt.Println(c.Diameter)
// 	fmt.Println(c.Area)
// 	fmt.Println(c.Perimeter)

// }

