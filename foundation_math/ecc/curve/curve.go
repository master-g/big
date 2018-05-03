package curve

import "math/big"

type Point struct {
	X *big.Int
	Y *big.Int
}

type Curve struct {
	A *big.Int // A y^2=x^3+Ax+B
	B *big.Int // B y^2=x^3+Ax+B
	P *big.Int // P is the prime number of finite field
	G *Point   // G generator point
	N *big.Int // N prime number of points in the group
}
