package geometry

type Curve struct {
	A int
	B int
}

func NewCurve(a, b int) *Curve {
	return &Curve{A: a, B: b}
}

func (c *Curve) Has(p *Point) bool {
	return p.Y*p.Y-p.X*p.X*p.X-c.A*p.X-c.B == 0
}

func (c *Curve) Add(p1, p2 *Point) *Point {
	if p1 == nil || p2 == nil || p1.X == p2.X {
		return nil
	}
	m := (p2.Y - p1.Y) / (p2.X - p1.X)
	b := p1.Y - 2*p1.X
}
