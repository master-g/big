package calc

import (
	"errors"
	"strings"
)

const (
	ExpressionCharSet = "0123456789+-x/%^"
)

var (
	// ErrorInvalidExpression invalid math expression
	ErrorInvalidExpression = errors.New("invalid expression")
)

// Expression contains information of a math expression
type Expression struct {
	origin string
}

// NewExpression parse math expression string and return Expression instance if it's valid
func NewExpression(s string) (*Expression, error) {
	if s == "" {
		return nil, ErrorInvalidExpression
	}

	s = strings.TrimSpace(s)
	for _, c := range s {
		if strings.Contains(ExpressionCharSet, c) {

		}
	}

	exp := &Expression{
		origin: s,
	}

	return exp, nil
}
