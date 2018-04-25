package main

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/master-g/big/src/foundation_math/ffcalc/calc"
)

var (
	ErrorInvalidFieldOrder = errors.New("invalid field order")
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ff [order] [exp] ...")
		os.Exit(1)
	}

	argOrder, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Printf("%v %v\n", ErrorInvalidFieldOrder, err)
		os.Exit(1)
	}
	order := big.NewInt(argOrder)
	fmt.Println(order)

	var expression string
	if len(os.Args) > 3 {
		expression = strings.Join(os.Args[2:], "")
	} else {
		expression = os.Args[2]
	}
	e, err := calc.NewExpression(expression)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(e)
}
