package sprint

// import "fmt"

func GenerateRange(min, max int) []int {
	// If min is bigger than max, return nil
	if min>= max {
		return nil
	}
	// Create parameter as a slice of integers with lenght based on the distance between min and max
	slice := make([]int, max-min)
	// We start from position 0
	position := 0
	// For each integer starting from min to max with step 1, the position of the slice is equal to the integer
	for n:=min; n<max; n++ {
		// for the first position we use the created position variable equal to 0
		slice[position] = n
		// then the position is moved forward of 1, so that in the next run of the for loop takes the position after the first
		position++
	}
	return slice
}

/* func main() {
	fmt.Println(GenerateRange(-1, 5))
} */