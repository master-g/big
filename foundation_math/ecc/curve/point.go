package curve

import (
	"fmt"
	"math/big"
)

type Point struct {
	X *big.Int
	Y *big.Int
}

func NewPoint(X, Y *big.Int) *Point {
	if X == nil {
		X = big.NewInt(0)
	}
	if Y == nil {
		Y = big.NewInt(0)
	}
	return &Point{
		X: X,
		Y: Y,
	}
}

func (p *Point) String() string {
	return fmt.Sprintf("(%v,%v)", p.X, p.Y)
}

func (p *Point) Equal(r *Point) bool {
	if p == r {
		return true
	}

	if p.X == r.X && p.Y == r.Y {
		return true
	}

	return p.X.Cmp(r.X) == 0 && p.Y.Cmp(r.Y) == 0
}

func (p *Point) Copy() *Point {
	x := big.NewInt(0).Set(p.X)
	y := big.NewInt(0).Set(p.Y)
	return NewPoint(x, y)
}
