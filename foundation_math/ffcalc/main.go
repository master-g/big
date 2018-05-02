package main

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/master-g/big/foundation_math/ffcalc/calc"
)

var (
	ErrorInvalidFieldOrder = errors.New("invalid field order")
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ffcalc prime [exp...]")
		os.Exit(1)
	}

	argPrime, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("invalid prime number")
		os.Exit(1)
	}
	prime := big.NewInt(int64(argPrime))
	if !prime.ProbablyPrime(64) {
		fmt.Printf("%v is probably not a prime number\n", argPrime)
		os.Exit(1)
	}

	var expression string
	if len(os.Args) > 3 {
		expression = strings.Join(os.Args[2:], "")
	} else {
		expression = os.Args[2]
	}
	rpn, err := calc.Parse(expression)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result, err := calc.Evaluate(rpn, prime)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}
