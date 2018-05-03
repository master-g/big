package geometry

import "fmt"

type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{X: x, Y: y}
}

func (p *Point) String() string {
	return fmt.Sprintf("(%v,%v)", p.X, p.Y)
}
