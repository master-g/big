package main

import (
	"fmt"

	"github.com/master-g/big/foundation_math/elliptic/geometry"
)

func main() {
	c := geometry.NewCurve(5, 7)
	fmt.Println(c.Has(geometry.NewPoint(2, 5)))
}
