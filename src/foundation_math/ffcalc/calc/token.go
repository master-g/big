package calc

import (
	"math/big"
)

// TokenType type of the token
type TokenType int

const (
	TokenParenthesisLeft  TokenType = 0
	TokenParenthesisRight TokenType = 1
	TokenLiteral          TokenType = 2
	TokenOperator         TokenType = 3
)

// Token structure
type Token struct {
	origin   string
	Type     TokenType
	Value    *big.Int
	Operator Operator
}

func (t *Token) IsParenthesis() bool {
	return t.Type == TokenParenthesisLeft || t.Type == TokenParenthesisRight
}

func (t *Token) IsOperator() bool {
	return t.Type == TokenOperator
}

func (t *Token) IsLiteral() bool {
	return t.Type == TokenLiteral
}

func (t *Token) String() string {
	return t.origin
}

type TokenStack []*Token

func NewTokenStack() *TokenStack {
	return &TokenStack{}
}

func (s TokenStack) Empty() bool {
	return len(s) == 0
}

func (s TokenStack) Peek() *Token {
	if len(s) == 0 {
		return nil
	}
	return s[len(s)-1]
}

func (s *TokenStack) Push(n *Token) {
	*s = append(*s, n)
}

func (s *TokenStack) Pop() *Token {
	if len(*s) == 0 {
		return nil
	}
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n
}
