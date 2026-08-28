package sprint
import "math"

func PointDistance(x1, y1, x2, y2 float64) float64 {
	dist_x := x2 - x1
	dist_y := y2 - y1
	dist_x_sq := dist_x * dist_x
	dist_y_sq := dist_y * dist_y
	sum_dist_sq := dist_x_sq + dist_y_sq
	return math.Sqrt(sum_dist_sq)
}