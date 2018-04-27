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

func (op Operator) Eval(operand1, operand2 *big.Int) *big.Int {
	if operand1 == nil || operand2 == nil {
		return nil
	}

	result := big.NewInt(0)

	switch op {
	case OperatorAdd:
		result.Add(operand1, operand2)
	case OperatorSub:
		result.Sub(operand1, operand2)
	case OperatorMul:
		result.Mul(operand1, operand2)
	case OperatorDiv:
		result.Div(operand1, operand2)
	case OperatorMod:
		result.Mod(operand1, operand2)
	case OperatorExp:
		result.Exp(operand1, operand2, nil)
	}

	return result
}
