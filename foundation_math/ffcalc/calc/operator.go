package calc

import (
	"math/big"
)

type Operator int

type Associativity int

const (
	OperatorAdd Operator = iota + 1
	OperatorSub
	OperatorMul
	OperatorDiv
	OperatorMod
	OperatorExp

	AssociativityLeft Associativity = iota + 1
	AssociativityRight
)

var (
	opAssocMap = map[Operator]Associativity{
		OperatorAdd: AssociativityLeft,
		OperatorSub: AssociativityLeft,
		OperatorMul: AssociativityLeft,
		OperatorDiv: AssociativityLeft,
		OperatorMod: AssociativityLeft,
		OperatorExp: AssociativityRight,
	}

	opPrecMap = map[Operator]int{
		OperatorAdd: 2,
		OperatorSub: 2,
		OperatorMul: 3,
		OperatorDiv: 3,
		OperatorMod: 3,
		OperatorExp: 4,
	}
)

func (op Operator) Associativity() Associativity {
	return opAssocMap[op]
}

func (op Operator) Precedence() int {
	return opPrecMap[op]
}

func mul(order, op1, op2 *big.Int) *big.Int {
	r := big.NewInt(0)
	r.Mul(op1, op2)
	// r.Mod(r, order)
	return r
}

func div(order, op1, op2 *big.Int) *big.Int {
	// a/b = a*b^-1 = a*b^(order-2)
	r := big.NewInt(0)
	r.Sub(order, big.NewInt(2))
	r.Exp(op2, r, nil)
	r.Mul(op1, r)
	// r.Mod(r, order)
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

func (op Operator) Eval(order, operand1, operand2 *big.Int) *big.Int {
	if order == nil || operand1 == nil || operand2 == nil {
		return nil
	}

	result := big.NewInt(0)

	switch op {
	case OperatorAdd:
		result.Add(operand1, operand2)
	case OperatorSub:
		result.Sub(operand1, operand2)
	case OperatorMul:
		return mul(order, operand1, operand2)
	case OperatorDiv:
		return div(order, operand1, operand2)
	case OperatorMod:
		result.Mod(operand1, operand2)
	case OperatorExp:
		return exp(order, operand1, operand2)
	}

	return result
}
