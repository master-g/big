package calc

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
