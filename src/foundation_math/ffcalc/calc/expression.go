package calc

import (
	"errors"
	"fmt"
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

	opTypeMap = map[string]Operator{
		"+": OperatorAdd,
		"-": OperatorSub,
		"x": OperatorMul,
		"/": OperatorDiv,
		"%": OperatorMod,
		"^": OperatorExp,
	}
)

func Parse(s string) error {
	tokens, err := tokenize(s)
	if err != nil {
		return err
	}

	ast, err := makeAbstractSyntaxTree(tokens)
	if err != nil {
		return err
	}

	fmt.Println(ast)

	return nil
}

func tokenize(s string) ([]*Token, error) {
	if s == "" {
		return nil, ErrorInvalidExpression
	}

	s = strings.TrimSpace(s)

	var sb strings.Builder
	tokenList := make([]*Token, 0)
	balance := 0
	for _, b := range s {
		c := string(b)
		token := &Token{}
		if unicode.IsDigit(b) {
			sb.WriteRune(b)
			continue
		} else if sb.Len() != 0 {
			token.origin = sb.String()
			v, _ := strconv.ParseInt(token.origin, 10, 64)
			token.Type = TokenLiteral
			token.Value = big.NewInt(v)
			tokenList = append(tokenList, token)
			sb.Reset()
		}

		token.origin = c

		if c == "(" {
			token.Type = TokenParenthesisLeft
			balance++
			continue
		} else if c == ")" {
			token.Type = TokenParenthesisRight
			balance--
			continue
		}

		if opType, ok := opTypeMap[c]; ok {
			token.Type = TokenOperator
			token.Operator = opType
		} else {
			return nil, ErrorInvalidCharacter
		}
	}

	if balance != 0 {
		return nil, ErrorMismatchedParenthesis
	}

	return tokenList, nil
}

func makeAbstractSyntaxTree(tokens []*Token) (*ASTNode, error) {
	outStack := NewASTStack()
	opStack := NewTokenStack()
	for _, t := range tokens {
		switch t.Type {
		case TokenLiteral:
			outStack.Push(NewASTNode(t, nil, nil))
		case TokenOperator:
			top := opStack.Peek()
			for top != nil && top.IsOperator() && (t.Operator.Associativity() == AssociativityLeft && t.Operator.Precedence() <= top.Operator.Precedence()) || (t.Operator.Associativity() == AssociativityRight && t.Operator.Precedence() < top.Operator.Precedence()) {
				outStack.AddNode(opStack.Pop())
				top = opStack.Peek()
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

	return outStack.Pop(), nil
}
