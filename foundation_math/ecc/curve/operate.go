package curve

import "math/big"

func add(op1, op2 *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Add(op1, op2)
	return r
}

func sub(op1, op2 *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Sub(op1, op2)
	return r
}

func mul(op1, op2 *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Mul(op1, op2)
	return r
}

func mod(order, op *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Mod(op, order)
	return r
}

func div(order, op1, op2 *big.Int) *big.Int {
	// a/b = a*b^-1 = a*b^(order-2)
	r := big.NewInt(0)
	r.Sub(order, big.NewInt(2))
	r.Exp(op2, r, nil)
	r.Mul(op1, r)
	return r
}

func exp(order, op1, op2 *big.Int) *big.Int {
	r := big.NewInt(0)
	if op2.Sign() < 0 {
		// a^-b = 1/a^b
		r.Abs(op2)
		r.Exp(op1, r, nil)
		return div(order, big.NewInt(1), r)
	} else {
		r.Exp(op1, op2, nil)
		return r
	}
}
