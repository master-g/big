package geometry

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{X: x, Y: y}
}
