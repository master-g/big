package calc

import "math/big"

var (
	zero *big.Int
	one  *big.Int
)

func init() {
	zero = big.NewInt(0)
	one = big.NewInt(1)
}

// Add return a + b
func Add(a, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, b)
}

// Sub return a - b
func Sub(a, b *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, b)
}

// Mod return a % b
func Mod(a, b *big.Int) *big.Int {
	return big.NewInt(0).Mod(a, b)
}

// Mul return op1 * op2
func Mul(op1, op2 *big.Int) *big.Int {
	return big.NewInt(0).Mul(op1, op2)
}

// Div return op1 / op2
func Div(order, op1, op2 *big.Int) *big.Int {
	invOp2 := Inv(order, op2)
	return big.NewInt(0).Mul(op1, invOp2)
}

// Inv return n^-1 mod order
func Inv(order, n *big.Int) *big.Int {
	gcd := big.NewInt(0)
	x := big.NewInt(0)
	y := big.NewInt(0)
	nx := big.NewInt(0)
	py := big.NewInt(0)
	gcd.GCD(x, y, n, order)
	nx.Mul(n, x)
	py.Mul(order, y)
	nx.Add(nx, py)
	nx.Mod(nx, order)
	if nx.Cmp(gcd) != 0 {
		return nil
	}
	if gcd.Cmp(one) != 0 {
		return nil
	}
	return x
}

// Exp return op1^op2
func Exp(order, op1, op2 *big.Int) *big.Int {
	r := big.NewInt(0)
	if op2.Sign() < 0 {
		// a^-b = 1/a^b
		r.Abs(op2)
		r.Exp(op1, r, nil)
		return Div(order, one, r)
	} else {
		r.Exp(op1, op2, nil)
		return r
	}
}
