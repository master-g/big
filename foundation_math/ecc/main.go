package main

import (
	"os"

	"fmt"
	"math/big"

	"github.com/master-g/big/foundation_math/ecc/curve"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	// cli
	app := &cli.App{
		Name:    "elliptic curve cryptography demo",
		Usage:   "ecc",
		Version: "0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "p",
				Usage: "prime number of the finite field",
			},
			&cli.StringFlag{
				Name:  "a",
				Usage: "y^2=x^3+ax+b",
			},
			&cli.StringFlag{
				Name:  "b",
				Usage: "y^2=x^3+ax+b",
			},
			&cli.StringFlag{
				Name:  "n",
				Usage: "prime number of points in the group",
			},
			&cli.StringFlag{
				Name:  "x",
				Usage: "x coordinate of the generator point",
			},
			&cli.StringFlag{
				Name:  "y",
				Usage: "y coordinate of the generator point",
			},
		},
		Action: func(c *cli.Context) error {

			a := big.NewInt(1)
			b := big.NewInt(1)
			p := big.NewInt(23)
			g := curve.NewPoint(big.NewInt(3), big.NewInt(10))
			v := curve.NewCurve(a, b, p, nil, g)
			P := curve.NewPoint(big.NewInt(3), big.NewInt(10))
			Q := curve.NewPoint(big.NewInt(9), big.NewInt(7))
			PN := v.InversePoint(P)
			PQ := v.Add(P, Q)
			P2 := v.Add(P, P)
			P27 := v.Mul(big.NewInt(27), P)
			fmt.Println(PN)
			fmt.Println(PQ)
			fmt.Println(P2)
			fmt.Println(P27)
			fmt.Println(v.N)
			return nil
		},
	}

	app.Run(os.Args)
}
