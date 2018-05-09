package calc

import "math/big"

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

// ExtendedEuclideanAlgorithm return a * x + b * y = gcd, where gcd is the greatest common divisor of a and b
func ExtendedEuclideanAlgorithm(a, b *big.Int) (gcd, x, y *big.Int) {
	s, oldS := big.NewInt(0), big.NewInt(1)
	t, oldT := big.NewInt(1), big.NewInt(0)
	r, oldR := big.NewInt(0).Set(b), big.NewInt(0).Set(a)
	for r.Sign() != 0 {
		quotient := big.NewInt(0).Div(oldR, r)
		tempR := big.NewInt(0).Set(quotient).Mul(quotient, r)
		tempR = tempR.Sub(oldR, tempR)
		oldR.Set(r)
		r.Set(tempR)

		tempS := big.NewInt(0).Set(quotient).Mul(quotient, s)
		tempS = tempS.Sub(oldS, tempS)
		oldS.Set(s)
		s.Set(tempS)

		tempT := big.NewInt(0).Set(quotient).Mul(quotient, t)
		tempT = tempT.Sub(oldT, tempT)
		oldT.Set(t)
		t.Set(tempT)
	}

	return oldR, oldS, oldT
}

// Mul return op1 * op2
func Mul(op1, op2 *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Mul(op1, op2)
	return r
}

// Div return op1 / op2
func Div(order, op1, op2 *big.Int) *big.Int {
	// a/b = a*b^-1 = a*b^(order-2)
	r := big.NewInt(0)
	r.Sub(order, big.NewInt(2))
	r.Exp(op2, r, nil)
	r.Mul(op1, r)
	// r.Mod(r, order)
	return r
}

// Inv return n^-1 mod order
func Inv(order, n *big.Int) *big.Int {
	gcd, x, y := ExtendedEuclideanAlgorithm(n, order)
	tmpNX := big.NewInt(0).Mul(n, x)
	tmpPY := big.NewInt(0).Mul(order, y)
	tmpNX.Add(tmpNX, tmpPY)
	tmpNX.Mod(tmpNX, order)
	if tmpNX.Cmp(gcd) != 0 {
		return nil
	}
	if gcd.Cmp(big.NewInt(1)) != 0 {
		return nil
	}
	return x.Mod(x, order)
}

// Exp return op1^op2
func Exp(order, op1, op2 *big.Int) *big.Int {
	r := big.NewInt(0)
	if op2.Sign() < 0 {
		// a^-b = 1/a^b
		r.Abs(op2)
		r.Exp(op1, r, nil)
		return Div(order, big.NewInt(1), r)
	} else {
		r.Exp(op1, op2, nil)
		return r
	}
}
