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

func (op Operator) Eval(order, operand1, operand2 *big.Int) *big.Int {
	if order == nil || operand1 == nil || operand2 == nil {
		return nil
	}

	result := big.NewInt(0)

	switch op {
	case OperatorAdd:
		return Add(operand1, operand2)
	case OperatorSub:
		return Sub(operand1, operand2)
	case OperatorMul:
		return Mul(operand1, operand2)
	case OperatorDiv:
		return Div(order, operand1, operand2)
	case OperatorMod:
		result.Mod(operand1, operand2)
	case OperatorExp:
		return Exp(order, operand1, operand2)
	}

	return result
}
