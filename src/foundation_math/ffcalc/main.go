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
	if len(os.Args) < 4 {
		fmt.Println("Usage: ffcalc [prime] [power] [exp] ...")
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

	argPower, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("invalid power number")
		os.Exit(1)
	}
	power := big.NewInt(int64(argPower))

	order := prime.Exp(prime, power, nil)

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
