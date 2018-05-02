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
	if len(os.Args) < 4 {
		fmt.Println("Usage: ffcalc prime degree [exp...]")
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

	argDegree, err := strconv.Atoi(os.Args[2])
	if err != nil || argDegree < 0 {
		fmt.Println("invalid degree number")
		os.Exit(1)
	}
	degree := big.NewInt(int64(argDegree))

	order := prime.Exp(prime, degree, nil)

	var expression string
	if len(os.Args) > 4 {
		expression = strings.Join(os.Args[3:], "")
	} else {
		expression = os.Args[3]
	}
	rpn, err := calc.Parse(expression)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result, err := calc.Evaluate(rpn, order)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}
