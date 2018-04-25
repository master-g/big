package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"

	"strings"

	"unicode"

	"github.com/btcsuite/goleveldb/leveldb/errors"
)

const (
	OperatorList = "+-x/%^"
)

var (
	ErrorInvalidExp = errors.New("invalid expression")
	ErrorInvalidOp  = fmt.Errorf("invalid operator, must be one of '%v'", OperatorList)
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ff [order] [exp]")
		os.Exit(1)
	}

	argOrder, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	order := big.NewInt(argOrder)

	exp := os.Args[2]
	var left, op, right string
	for i, c := range exp {
		if !unicode.IsDigit(c) && i > 0 && i < len(exp)-1 {
			left = exp[:i]
			op = exp[i : i+1]
			right = exp[i+1:]
			break
		}
	}
	if left == "" || op == "" || right == "" {
		fmt.Println(ErrorInvalidExp)
		os.Exit(1)
	}

	argLeft, err := strconv.ParseInt(left, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	operand1 := big.NewInt(argLeft)

	if len(op) != 1 || !strings.ContainsAny(OperatorList, op) {
		fmt.Println(ErrorInvalidOp)
		os.Exit(1)
	}

	argRight, err := strconv.ParseInt(right, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	operand2 := big.NewInt(argRight)

	z := big.NewInt(0)
	switch op {
	case "+":
		z.Add(operand1, operand2)
	case "-":
		z.Sub(operand1, operand2)
	case "x":
		z.Mul(operand1, operand2)
	case "/":
		z.Div(operand1, operand2)
	case "%":
		z.Mod(operand1, operand2)
	case "^":
		z.Exp(operand1, operand2, order)
	}

	r := big.NewInt(0)
	r.Mod(z, order)

	fmt.Println(r)
}
