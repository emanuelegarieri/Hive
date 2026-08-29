//go:build ignore
// +build ignore

package main


type Contestant struct {
	Name 		string
	Scores 	 	[]int
}


func GetWinner(c1, c2 Contestant) string {

	var sum1 int
	var sum2 int

	for r := range len(c1.Scores) {
		sum1 = sum1 + c1.Scores[r]
	}
	for r := range len(c2.Scores) {
		sum2 = sum2 + c2.Scores[r]
	}

	if sum1 < sum2 {
		return c2.Name
	}
	 return c1.Name
}



