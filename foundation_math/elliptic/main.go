package main

import (
	"fmt"

	"github.com/master-g/big/foundation_math/elliptic/geometry"
)

func main() {
	c := geometry.NewCurve(5, 7)
	p1 := geometry.NewPoint(3, 7)
	p2 := geometry.NewPoint(18, 77)
	fmt.Println(c.Add(p1, p2))
}
