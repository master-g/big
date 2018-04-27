package calc

import (
	"errors"
	"math/big"
	"strconv"
	"strings"
	"unicode"
)

var (
	// ErrorInvalidExpression invalid math expression
	ErrorInvalidExpression = errors.New("invalid expression")

	// ErrorInvalidCharacter invalid math expression
	ErrorInvalidCharacter = errors.New("invalid character")

	// ErrorMismatchedParenthesis indicates there are unmatched parenthesis
	ErrorMismatchedParenthesis = errors.New("invalid mismatched parenthesis")

	// ErrorInvalidRPN invalid reverse polish notation
	ErrorInvalidRPN = errors.New("invalid RPN")

	opTypeMap = map[string]Operator{
		"+": OperatorAdd,
		"-": OperatorSub,
		"*": OperatorMul,
		"/": OperatorDiv,
		"%": OperatorMod,
		"^": OperatorExp,
	}
)

// Parse expression to reverse polish notation
func Parse(s string) ([]*Token, error) {
	tokens, err := tokenize(s)
	if err != nil {
		return nil, err
	}

	ast := makeAbstractSyntaxTree(tokens)
	if ast == nil {
		return nil, ErrorInvalidExpression
	}

	visitor := NewCalcVisitor()
	ast.Accept(visitor)

	return visitor.RPN, nil
}

// Evaluate the arithmetic value of a reverse polish notation
func Evaluate(tokens []*Token, order *big.Int) (*big.Int, error) {
	if len(tokens) == 0 {
		return nil, ErrorInvalidRPN
	}

	stack := NewTokenStack()
	for _, t := range tokens {
		if t.IsOperator() {
			token2 := stack.Pop()
			token1 := stack.Pop()
			if token1.Type != TokenLiteral || token1.Value == nil || token2.Type != TokenLiteral || token2.Value == nil {
				return nil, ErrorInvalidRPN
			}
			result := t.Operator.Eval(token1.Value, token2.Value)
			result.Mod(result, order)
			stack.Push(&Token{
				origin: result.String(),
				Type:   TokenLiteral,
				Value:  result,
			})
		} else {
			stack.Push(t)
		}
	}

	value := stack.Pop()
	if !stack.Empty() {
		return nil, ErrorInvalidRPN
	}

	if value.Type != TokenLiteral {
		return nil, ErrorInvalidRPN
	}

	return value.Value, nil
}

func preprocess(s string) string {
	var sb strings.Builder
	// remove all spaces
	s = strings.TrimSpace(s)

	// process parenthesis
	s = strings.Replace(s, ")(", ")*(", -1)
	for i, b := range s {
		c := string(b)
		if c == "(" && i != 0 && unicode.IsDigit([]rune(s)[i-1]) {
			sb.WriteString("*")
		}
		sb.WriteRune(b)
	}
	s = sb.String()

	// process negative numbers
	for i, b := range s {
		c := string(b)
		if c == "-" {
			// TODO:
			// 1--2 -> 1-(0-2) 			a. operator before '-'
			// 2*-(3+4) -> 2*(0-(3+4)) 	b. '(' after '-'
		}
	}

	return s
}

func tokenize(s string) ([]*Token, error) {
	if s == "" {
		return nil, ErrorInvalidExpression
	}

	// process space, parenthesis and negative numbers
	s = preprocess(s)

	// tokenize
	var sb strings.Builder
	tokenList := make([]*Token, 0)
	balance := 0
	for i, b := range s {
		c := string(b)
		if unicode.IsDigit(b) {
			sb.WriteRune(b)
			if i < len(s)-1 {
				continue
			}
		}

		if sb.Len() != 0 {
			origin := sb.String()
			v, _ := strconv.ParseInt(origin, 10, 64)
			token := &Token{
				origin: origin,
				Type:   TokenLiteral,
				Value:  big.NewInt(v),
			}
			tokenList = append(tokenList, token)
			sb.Reset()

			if unicode.IsDigit(b) {
				continue
			}
		}

		token := &Token{origin: c}
		if c == "(" {
			token.Type = TokenParenthesisLeft
			balance++
			tokenList = append(tokenList, token)
			continue
		} else if c == ")" {
			token.Type = TokenParenthesisRight
			balance--
			tokenList = append(tokenList, token)
			continue
		}

		if opType, ok := opTypeMap[c]; ok {
			token.Type = TokenOperator
			token.Operator = opType
			tokenList = append(tokenList, token)
		} else {
			return nil, ErrorInvalidCharacter
		}
	}

	if balance != 0 {
		return nil, ErrorMismatchedParenthesis
	}

	return tokenList, nil
}

func makeAbstractSyntaxTree(tokens []*Token) *ASTNode {
	outStack := NewASTStack()
	opStack := NewTokenStack()
	for _, t := range tokens {
		switch t.Type {
		case TokenLiteral:
			outStack.Push(NewASTNode(t, nil, nil))
		case TokenOperator:
			var top *Token
			for {
				if opStack.Empty() {
					break
				}
				top = opStack.Peek()
				if top.IsOperator() && (t.Operator.Associativity() == AssociativityLeft && t.Operator.Precedence() <= top.Operator.Precedence()) || (t.Operator.Associativity() == AssociativityRight && t.Operator.Precedence() < top.Operator.Precedence()) {
					outStack.AddNode(opStack.Pop())
					top = opStack.Peek()
				}
			}
			opStack.Push(t)
		case TokenParenthesisLeft:
			opStack.Push(t)
		case TokenParenthesisRight:
			top := opStack.Peek()
			for top != nil && top.Type != TokenParenthesisLeft {
				outStack.AddNode(opStack.Pop())
				top = opStack.Peek()
			}
			opStack.Pop()
		}
	}

	for opStack.Peek() != nil {
		outStack.AddNode(opStack.Pop())
	}

	return outStack.Pop()
}

// RPNVisitor
type RPNVisitor struct {
	ASTNodeVisitor
	RPN []*Token
}

// NewCalcVisitor returns new RPNVisitor
func NewCalcVisitor() *RPNVisitor {
	return &RPNVisitor{
		RPN: make([]*Token, 0),
	}
}

// Visit interface
func (dv *RPNVisitor) Visit(n *ASTNode) {
	dv.RPN = append([]*Token{n.Token}, dv.RPN...)
}
