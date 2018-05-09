package curve

import (
	"math/big"
)

type Curve struct {
	A *big.Int // A y^2=x^3+Ax+B
	B *big.Int // B y^2=x^3+Ax+B
	P *big.Int // P is the prime number of finite field
	G *Point   // G generator point
	N *big.Int // N prime number of points in the group
}

func NewCurve(A, B, P, N *big.Int, G *Point) *Curve {
	if P == nil || G == nil {
		return nil
	}
	if A == nil && B == nil {
		return nil
	}
	if A == nil {
		A = big.NewInt(0)
	}
	if B == nil {
		B = big.NewInt(0)
	}
	curve := &Curve{
		A: A,
		B: B,
		P: P,
		N: N,
		G: G,
	}

	if !curve.Has(G) {
		return nil
	}

	if curve.N == nil {
		curve.N = curve.findN(curve.G, big.NewInt(512))
	}

	return curve
}

// Has returns true is p is on curve c (y^2=x^3+ax+b)
func (c *Curve) Has(p *Point) bool {
	y2 := exp(c.P, p.Y, big.NewInt(2))
	y2.Mod(y2, c.P)
	x3 := exp(c.P, p.X, big.NewInt(3))
	ax := mul(p.X, c.A)
	r := big.NewInt(0).Add(x3, ax)
	r.Add(r, c.B)
	r.Mod(r, c.P)
	r.Sub(r, y2)
	return r.Sign() == 0
}

func (c *Curve) InversePoint(p *Point) *Point {
	y := big.NewInt(0).Neg(p.Y)
	y.Mod(y, c.P)
	x := big.NewInt(0).Set(p.X)
	return NewPoint(x, y)
}

func (c *Curve) Add(p1, p2 *Point) *Point {
	k := big.NewInt(0)
	x3 := big.NewInt(0)
	y3 := big.NewInt(0)
	if p1.Equal(p2) {
		// k = (3x2+a)/2y1
		dy := exp(c.P, p2.X, big.NewInt(2))
		dy.Mul(dy, big.NewInt(3))
		dy.Add(dy, c.A)
		dx := mul(p1.Y, big.NewInt(2))
		k = div(c.P, dy, dx)
	} else {
		// k = (y2-y1)/(x2-x1)
		dy := big.NewInt(0).Sub(p2.Y, p1.Y)
		dx := big.NewInt(0).Sub(p2.X, p1.X)
		k = div(c.P, dy, dx)
	}

	x3.Exp(k, big.NewInt(2), nil)
	x3.Sub(x3, p1.X)
	x3.Sub(x3, p2.X)
	x3.Mod(x3, c.P)
	y3.Sub(p1.X, x3)
	y3.Mul(y3, k)
	y3.Sub(y3, p1.Y)
	y3.Mod(y3, c.P)

	return NewPoint(x3, y3)
}

func bits(n *big.Int) []uint {
	bits := make([]uint, n.BitLen())
	for i := 0; i < n.BitLen(); i++ {
		bits[i] = n.Bit(i)
	}
	return bits
}

func (c *Curve) Mul(n *big.Int, p *Point) *Point {
	if n.Sign() <= 0 {
		return nil
	}
	if n.Cmp(big.NewInt(1)) == 0 {
		return p.Copy()
	}
	bits := bits(n)
	np := p.Copy()
	var result *Point
	for i := len(bits) - 1; i >= 0; i-- {
		if bits[i] == 1 {
			if result == nil {
				result = np.Copy()
			} else {
				result = c.Add(result, np)
			}
		}
		np = c.Add(np, np)
	}

	return result
}

func (c *Curve) findN(p *Point, tries *big.Int) *big.Int {
	if p == nil {
		return nil
	}

	np := c.InversePoint(p)
	cp := p.Copy()
	one := big.NewInt(1)
	i := big.NewInt(1)
	limit := big.NewInt(0)
	for {
		// i++
		i.Add(i, one)
		cp = c.Add(cp, p)
		if cp.Equal(np) {
			return i
		}
		if tries != nil && limit.Sub(tries, i).Sign() <= 0 {
			break
		}
	}
	return nil
}
