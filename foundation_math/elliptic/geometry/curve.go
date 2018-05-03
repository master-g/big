package geometry

type Curve struct {
	A float64
	B float64
}

func NewCurve(a, b float64) *Curve {
	return &Curve{A: a, B: b}
}

func (c *Curve) Has(p *Point) bool {
	return p.Y*p.Y-p.X*p.X*p.X-c.A*p.X-c.B == 0
}

func (c *Curve) Add(p1, p2 *Point) *Point {
	if p1 == nil || p2 == nil || p1.X == p2.X {
		return nil
	}
	if !c.Has(p1) || !c.Has(p2) {
		return nil
	}
	m := (p1.Y - p2.Y) / (p1.X - p2.X)
	v := p1.Y - m*p1.X

	x3 := m*m - p1.X - p2.X
	y3 := m*x3 + v
	return NewPoint(x3, -y3)
}
